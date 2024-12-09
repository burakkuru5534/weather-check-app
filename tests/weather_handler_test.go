package handler_test

import (
	"github.com/stretchr/testify/mock"
	_ "weather-check-app/services"
)

type MockWeatherService struct {
	mock.Mock
}

func (m *MockWeatherService) FetchWeatherAPI(location string) (float64, error) {
	args := m.Called(location)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockWeatherService) FetchWeatherStack(location string) (float64, error) {
	args := m.Called(location)
	return args.Get(0).(float64), args.Error(1)
}
