package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"

	"github/mars-terminal/telegram-notification.com/internal/server/telegram"
	"github/mars-terminal/telegram-notification.com/internal/service/weather"
)

func main() {
	log := logrus.WithField("package", "main")

	if err := initConfig(); err != nil {
		log.Fatal("could not connect to config")
		return
	}

	var opts = struct {
		ApiKey string
		Token  string
	}{
		ApiKey: viper.GetString("api_key"),
		Token:  viper.GetString("token"),
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  opts.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("could not get the token")
		return
	}

	service := weather.NewService(opts.ApiKey)
	telegram.NewHandler(bot, service).InitHandler()
	log.Info("[INFO] - Started")

	bot.Start()
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
