package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (r *BirthdayRouter) PickCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "list":
		r.list(*update.Message)
	case "soon":
		r.soon(*update.Message)
	case "setDate":
		r.setDate(update)
	case "help":
		r.help(*update.Message)
	case "delete":
		r.delete(*update.Message)
	}
}
