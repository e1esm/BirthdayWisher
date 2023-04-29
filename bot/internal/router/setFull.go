package router

import (
	"BirthdayWisherBot/internal/models/bridge"
	"BirthdayWisherBot/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"log"
	"regexp"
	"strings"
)

func (r *BirthdayRouter) setFull(message tgbotapi.Message) {
	regex := regexp.MustCompile("^\\s*(3[01]|[12][0-9]|0?[1-9])\\.(1[012]|0?[1-9])\\.((?:19|20)\\d{2})\\s*$")
	if !regex.MatchString(message.CommandArguments()) {
		log.Println(message.CommandArguments())
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Invalid date: %v", message.CommandArguments()))
		r.bot.Send(msg)
		return
	}
	splittedMessage := strings.Split(message.CommandArguments(), ".")
	date := fmt.Sprintf("%s-%s-%s", splittedMessage[2], splittedMessage[1], splittedMessage[0])
	chat := bridge.NewChat(message.Chat.ID)
	user := bridge.NewUser(message.From.ID, date, *chat, message.From.FirstName+message.From.LastName)

	err := r.ConnectorService.SaveUser(*user)
	if err != nil {
		utils.Logger.Error("Failed to save user's info into DB", zap.String("user", user.Username))
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Не получилось сохранить данные %s в БД", message.From.FirstName))
		r.bot.Send(msg)
		return
	}
	utils.Logger.Info("Successfully added user's info into DB", zap.String("user", user.Username))
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Данные %s были добавлены в БД", message.From.FirstName))
	r.bot.Send(msg)

}
