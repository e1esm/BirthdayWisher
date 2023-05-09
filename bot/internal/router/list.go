package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

//TODO: list all the birthdays of chat members

func (r *BirthdayRouter) list(message tgbotapi.Message) {
	pdf, _ := r.ConnectorService.GetChatStatistics(message.Chat.ID)
	doc := tgbotapi.NewDocument(message.Chat.ID, pdf)
	r.bot.Send(doc)
}
