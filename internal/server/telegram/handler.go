package telegram

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"

	"github/mars-terminal/telegram-notification.com/internal/service"
)

var logger = logrus.WithField("service", "weather")

type Handler struct {
	bot *telebot.Bot

	service service.Service
}

func NewHandler(bot *telebot.Bot, service service.Service) *Handler {
	return &Handler{bot: bot, service: service}
}

func (h *Handler) InitHandler() {
	h.bot.Handle(telebot.OnText, h.weatherNotification)
}
