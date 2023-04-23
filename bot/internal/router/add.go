package router

import (
	"BirthdayWisherBot/internal/models/bridge"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
	"strings"
)

// TODO add changing bool argument for call from change/pickCommand
func (r *BirthdayRouter) add(message tgbotapi.Message) {
	regex := regexp.MustCompile("^(?:0[1-9]|[12][0-9]|3[01]).(?:0[1-9]|1[012])$")
	if !regex.MatchString(message.CommandArguments()) {
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Invalid date: %s", message.CommandArguments()))
		r.bot.Send(msg)
		return
	}
	splittedMessage := strings.Split(message.CommandArguments(), ".")
	date := fmt.Sprintf("1970-%s-%s", splittedMessage[1], splittedMessage[0])
	chat := bridge.NewChat(message.Chat.ID)
	user := bridge.NewUser(message.From.ID, date, *chat)
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("%#+v", user))
	r.bot.Send(msg)
}
