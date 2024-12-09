package utils

import (
	"errors"
	"sync"
	"time"
	"weather-check-app/db"
	"weather-check-app/models"
	"weather-check-app/services"
)

var requestGroups = make(map[string]*groupInfo)
var mu sync.Mutex

type groupInfo struct {
	clients   []chan float64
	timer     *time.Timer
	requested bool
	mu        sync.Mutex
}

func GroupedWeatherQuery(location string) (float64, error) {
	mu.Lock()
	group, exists := requestGroups[location]
	if !exists {
		group = &groupInfo{
			clients:   []chan float64{},
			timer:     time.NewTimer(5 * time.Second),
			requested: false,
		}
		requestGroups[location] = group
		go processGroup(location, group)
	}
	mu.Unlock()

	clientChan := make(chan float64, 1)
	group.mu.Lock()
	group.clients = append(group.clients, clientChan)
	if len(group.clients) >= 10 && !group.requested {
		group.requested = true
		group.timer.Stop()
		go processGroup(location, group)
	}
	group.mu.Unlock()

	return <-clientChan, nil
}

func processGroup(location string, group *groupInfo) {
	<-group.timer.C
	group.mu.Lock()
	if group.requested {
		group.mu.Unlock()
		return
	}
	group.requested = true
	group.mu.Unlock()

	temp1, temp2, err := fetchWeatherData(location)
	group.mu.Lock()
	for _, client := range group.clients {
		if err != nil {
			client <- -1
		} else {
			client <- (temp1 + temp2) / 2
		}
		close(client)
	}
	group.mu.Unlock()
}

func fetchWeatherData(location string) (float64, float64, error) {
	var wg sync.WaitGroup
	var temp1, temp2 float64
	var err1, err2 error
	wg.Add(2)

	// Fetch temperature from WeatherAPI
	go func() {
		defer wg.Done()
		temp1, err1 = services.FetchWeatherAPI(location)
	}()

	// Fetch temperature from WeatherStack
	go func() {
		defer wg.Done()
		temp2, err2 = services.FetchWeatherStack(location)
	}()

	wg.Wait()

	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("failed to fetch weather data")
	}

	// Log the query into the database
	db.DB.Create(&models.WeatherQuery{
		Location:            location,
		Service1Temperature: temp1,
		Service2Temperature: temp2,
		RequestCount:        len(requestGroups[location].clients),
		CreatedAt:           time.Now(),
	})

	return temp1, temp2, nil
}
