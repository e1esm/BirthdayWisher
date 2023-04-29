package router

import (
	"BirthdayWisherBot/internal/models/bridge"
	"BirthdayWisherBot/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

func (r *BirthdayRouter) set(message tgbotapi.Message) {
	regex := regexp.MustCompile("^(?:0[1-9]|[12][0-9]|3[01]).(?:0[1-9]|1[012])$")
	if !regex.MatchString(message.CommandArguments()) {
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Invalid date: %s", message.CommandArguments()))
		r.bot.Send(msg)
		return
	}
	splittedMessage := strings.Split(message.CommandArguments(), ".")
	date := fmt.Sprintf("1970-%s-%s", splittedMessage[1], splittedMessage[0])
	chat := bridge.NewChat(message.Chat.ID)
	user := bridge.NewUser(message.From.ID, date, *chat, message.From.FirstName+message.From.LastName)

	err := r.ConnectorService.SaveUser(*user)
	if err != nil {
		utils.Logger.Error("Failed to save user's info into DB", zap.String("user", user.Username))
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Не получилось сохранить данные %s в БД", message.From.FirstName))
		r.bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Данные %s были внесены в БД", message.From.FirstName))
	utils.Logger.Info("Successfully added user's info into DB", zap.String("user", user.Username))
	r.bot.Send(msg)
}
