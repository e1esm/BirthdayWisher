package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

// TODO list all the members that're going to have a birthday this month
func (r *BirthdayRouter) soon(message tgbotapi.Message) {
	response, _ := r.ConnectorService.GetSoonBirthdays(message.Chat.ID)

	sb := strings.Builder{}
	sb.WriteString("=====\n*Предстоящие дни рождения этого месяца*\n")
	for _, birthday := range response.SoonBirthdays {
		sb.WriteString(birthday.Username)
		sb.WriteString(" : ")
		sb.WriteString(birthday.BirthdayDate)
		sb.WriteString("\n\n")
	}
	sb.WriteString("=====\n")

	msg := tgbotapi.NewMessage(message.Chat.ID, sb.String())
	msg.ParseMode = "markdown"
	r.bot.Send(msg)
}
