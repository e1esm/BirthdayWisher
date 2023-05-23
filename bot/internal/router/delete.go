package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (r *BirthdayRouter) delete(message tgbotapi.Message) {
	err := r.ConnectorService.DeleteUser(message.From.ID, message.Chat.ID)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Не удалось удалить ваши данные")
		r.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Ваши данные успешно удалены")
	r.bot.Send(msg)
}
