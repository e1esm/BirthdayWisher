package router

import (
	"BirthdayWisherBot/internal/models/bridge"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"regexp"
	"strings"
)

// TODO add changing bool argument for call from change/pickCommand
func (r *BirthdayRouter) addFull(message tgbotapi.Message) {
	regex := regexp.MustCompile("(0[1-9]|[1-2][0-9]|3[0-1])\\.(0[1-9]|1[0-2])\\.\\d{4}$")
	if !regex.MatchString(message.CommandArguments()) {
		log.Println(message.CommandArguments())
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Invalid date: %v", message.CommandArguments()))
		r.bot.Send(msg)
		return
	}
	splittedMessage := strings.Split(message.CommandArguments(), ".")
	date := fmt.Sprintf("%s-%s-%s", splittedMessage[2], splittedMessage[1], splittedMessage[0])
	chat := bridge.NewChat(message.Chat.ID)
	user := bridge.NewUser(message.From.ID, date, *chat)

	err := r.bridgeService.SaveUser(*user)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Не получилось сохранить данные %s в БД", message.From.FirstName))
		r.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Данные %s были добавлены в БД", message.From.FirstName))
	r.bot.Send(msg)

}
