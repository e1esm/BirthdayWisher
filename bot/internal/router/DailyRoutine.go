package router

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r *BirthdayRouter) DailyChecker() {

	r.scheduler.Every(1).Day().At("00:00").Do(func() {
		responses, err := r.ConnectorService.DailyRetriever()
		if err != nil {
			for _, response := range responses {
				for _, chat := range response.ChatIDs {
					message := tgbotapi.NewMessage(chat, fmt.Sprintf("Не получилось найти никакую информацию о пользователе %s", response.Username))
					r.bot.Send(message)
				}
			}
		}
		for _, response := range responses {
			for _, chat := range response.ChatIDs {
				message := tgbotapi.NewMessage(chat, response.CongratulationSentence)
				r.bot.Send(message)
			}
		}
	})
}
