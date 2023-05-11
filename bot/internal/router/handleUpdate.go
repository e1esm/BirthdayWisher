package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (r *BirthdayRouter) HandleUpdate(update tgbotapi.Update) {

	if update.CallbackQuery != nil {
		r.handleCallback(update)
	}

	if update.Message != nil {
		if update.Message.IsCommand() {
			r.PickCommand(update)
		}
	}
}
