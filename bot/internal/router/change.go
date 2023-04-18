package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (r *BirthdayRouter) change(message tgbotapi.Message) {
	splittedMessage := strings.Split(message.CommandArguments(), ".")
	if len(splittedMessage) == 3 {
		r.addFull(message)
	}
	r.add(message)
}
