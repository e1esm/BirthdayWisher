package router

import (
	"BirthdayWisherBot/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (r *BirthdayRouter) DailyBirthdayChecker() {
	utils.Logger.Info("Started checking for today's birthdays")
	responses, err := r.ConnectorService.DailyRetriever()
	if err != nil {
		utils.Logger.Error("Couldn't have retrieved any info from ConnectorService")
	}
	for _, response := range responses {
		for _, chat := range response.ChatIDs {
			if err == nil {
				utils.Logger.Info("Congratulated a user", zap.String("username", response.Username))
				message := tgbotapi.NewMessage(chat.ChatID, fmt.Sprintf("[%s](tg://user?id=%d),\n%s", response.Username, response.UserID, response.CongratulationSentence))
				message.ParseMode = "markdown"
				r.bot.Send(message)
			} else {
				utils.Logger.Error("No info provided for this user", zap.String("user", response.Username))
				message := tgbotapi.NewMessage(chat.ChatID, fmt.Sprintf("Не получилось найти никакую информацию о пользователе %s", response.Username))
				r.bot.Send(message)
			}

		}
	}
}
