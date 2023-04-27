package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// TODO list all the members that're going to have a birthday this month
func (r *BirthdayRouter) soon(message tgbotapi.Message) {
	response, _ := r.ConnectorService.GetSoonBirthdays(message.Chat.ID)
	msg := tgbotapi.NewMessage(message.Chat.ID, response.String())
	r.bot.Send(msg)
}
