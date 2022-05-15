package weather

import (
	"fmt"
	"io"
	"net/http"

	"github/mars-terminal/telegram-notification.com/internal/entities"
)

const url = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"

func (s *Service) GetWeather(city string) (*entities.ApiResponse, error) {
	log := logger.WithField("read", "get weather")

	response, err := http.Get(fmt.Sprintf(url, city, s.ApiKey))
	if err != nil {
		log.WithError(err).Error("can't make request to call open weather api")
		return nil, err
	}

	read, err := io.ReadAll(response.Body)
	if err != nil {
		log.WithError(err).Error("can't read request body after call open weather api")
		return nil, err
	}

	resp := &entities.ApiResponse{}

	if err := resp.Unmarshal(read); err != nil {
		log.WithError(err).Error("can't unmarshal request body after read open weather api data")
		return nil, err
	}

	if resp.Message != "" {
		log.Error(resp.Message)
		return nil, fmt.Errorf(resp.Message)
	}

	return &entities.ApiResponse{
		Main:     resp.Main,
		Timezone: resp.Timezone / 3600,
		Wind:     resp.Wind,
		Name:     resp.Name,
	}, nil

}
