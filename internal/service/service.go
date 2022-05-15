package service

import "github/mars-terminal/telegram-notification.com/internal/entities"

type Service interface {
	GetWeather(city string) (*entities.ApiResponse, error)
}
