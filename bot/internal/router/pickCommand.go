package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (r *BirthdayRouter) PickCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "change":
		r.change(*update.Message)
	case "list":
		r.list(*update.Message)
	case "soon":
		r.soon(*update.Message)
	case "addDate":
		r.addDate(update)
	case "help":
		r.help(*update.Message)
	}
}
