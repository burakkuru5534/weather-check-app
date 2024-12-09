package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherStackResponse struct {
	Current struct {
		Temperature float64 `json:"temperature"`
	} `json:"current"`
}

func FetchWeatherStack(location string) (float64, error) {
	apiKey := "838c0d5e8fcc1dbbc66e8c1c0a14c6e5"
	url := fmt.Sprintf("http://api.weatherstack.com/current?access_key=%s&query=%s", apiKey, location)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var response WeatherStackResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}

	return response.Current.Temperature, nil
}
