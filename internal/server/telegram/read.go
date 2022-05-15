package telegram

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

func (h *Handler) weatherNotification(c telebot.Context) error {

	log := logger.WithField("telegram", "weather notification")

	city := c.Message().Text

	data, err := h.service.GetWeather(city)
	if err != nil {
		log.WithError(err).Error("could not get data")
		return c.Send(err.Error())
	}

	byte := fmt.Sprintf(`
City: %s,
Temperature: %v °C,
Wind speed: %v m/s,
TimeZone: %v,
`, data.Name, data.Main.Temp, data.Wind.Speed, data.Timezone)

	_, err = h.bot.Send(c.Sender(), byte)
	if err != nil {
		log.WithError(err).Error("could not get data")
		return err
	}
	return nil
}
