package weather

import "github.com/sirupsen/logrus"

type Service struct {
	ApiKey string
}

func NewService(apiKey string) *Service {
	return &Service{ApiKey: apiKey}
}

var logger = logrus.WithField("service", "weather")
