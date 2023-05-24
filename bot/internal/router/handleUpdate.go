package router

import (
	"BirthdayWisherBot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r *BirthdayRouter) HandleUpdate(update tgbotapi.Update) {

	if update.CallbackQuery != nil {
		r.handleCallback(update)
	}

	if update.Message != nil {
		if update.Message.IsCommand() {
			r.PickCommand(update)
		}
		if user := update.Message.LeftChatMember; user != nil {
			err := r.ConnectorService.DeleteUser(user.ID, update.FromChat().ID)
			if err != nil {
				utils.Logger.Error("Couldn't have delete a left user")
				return
			}
		}
	}
}
