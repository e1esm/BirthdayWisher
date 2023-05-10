package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"time"
)

func (r *BirthdayRouter) soon(message tgbotapi.Message) {
	response, _ := r.ConnectorService.GetSoonBirthdays(message.Chat.ID)

	sb := strings.Builder{}
	sb.WriteString("* Предстоящие дни рождения этого месяца * \n-------\n")
	for i, birthday := range response.SoonBirthdays {
		sb.WriteString(birthday.Username)
		sb.WriteString(" : ")
		birthdayDate := strconv.Itoa(time.Now().Year()) + birthday.BirthdayDate[4:10]
		sb.WriteString(birthdayDate)
		if i == len(response.SoonBirthdays)-1 {
			sb.WriteString("\n")
			break
		}
		sb.WriteString("\n\n")
	}
	sb.WriteString("-------\n")

	msg := tgbotapi.NewMessage(message.Chat.ID, sb.String())
	msg.ParseMode = "markdown"
	_, err := r.bot.Send(msg)
	if err != nil {
		msg.ParseMode = ""
		r.bot.Send(msg)
	}
}
