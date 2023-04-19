package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (r *BirthdayRouter) PickCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "add":
		r.add(*update.Message)
	case "addFull":
		r.add(*update.Message)
	case "change":
		r.change(*update.Message)
	case "list":
	case "soon":
	case "help":
		r.help(*update.Message)
	}
}
