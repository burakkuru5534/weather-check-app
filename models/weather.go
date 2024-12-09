package models

import "time"

type WeatherQuery struct {
	ID                  uint   `gorm:"primaryKey"`
	Location            string `gorm:"index"`
	Service1Temperature float64
	Service2Temperature float64
	RequestCount        int
	CreatedAt           time.Time
}
