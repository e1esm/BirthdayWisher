package router

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (r *BirthdayRouter) DailyBirthdayChecker() {
	responses, err := r.ConnectorService.DailyRetriever()
	if err != nil {
		log.Println(err)
	}
	for _, response := range responses {
		for _, chat := range response.ChatIDs {
			if err == nil {
				message := tgbotapi.NewMessage(chat.ChatID, fmt.Sprintf("[%s](tg://user?id=%d),\n%s", response.Username, response.UserID, response.CongratulationSentence))
				message.ParseMode = "markdown"
				r.bot.Send(message)
			} else {
				message := tgbotapi.NewMessage(chat.ChatID, fmt.Sprintf("Не получилось найти никакую информацию о пользователе %s", response.Username))
				r.bot.Send(message)
			}

		}
	}
}
