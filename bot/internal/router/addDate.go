package router

import (
	state "BirthdayWisherBot/utils"
	"fmt"
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"time"
)

func (r *BirthdayRouter) setDate(update tgbotapi.Update) {
	currentStateConfig := state.StateConfig{
		CurrentState: state.START,
		UserID:       update.SentFrom().ID,
		ChatID:       update.FromChat().ID,
		MessageID:    update.Message.MessageID,
		Offset:       0,
	}
	state.RWMapInstance.UpdateConfig(currentStateConfig)
	r.addYear(update)
}

func (r *BirthdayRouter) addYear(update tgbotapi.Update) {

	v, _ := state.RWMapInstance.GetConfig(update.SentFrom().ID)
	v.CurrentState = state.YEAR
	state.RWMapInstance.UpdateConfig(v)

	arrRows := make([][]tgbotapi.InlineKeyboardButton, 5)
	currentOffset := v.Offset

	currentYear := time.Now().Year()
	for i := 0; i < 4; i++ {
		arrRows[i] = make([]tgbotapi.InlineKeyboardButton, 3)
		for j := 0; j < 3; j++ {
			year := fmt.Sprintf("%d", currentYear-currentOffset)
			arrRows[i][j] = tgbotapi.NewInlineKeyboardButtonData(year, year)
			currentOffset++
		}
	}
	arrRows[len(arrRows)-1] = []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData(emoji.LeftArrow.String(), "-12"),
		tgbotapi.NewInlineKeyboardButtonData(emoji.CrossMark.String(), fmt.Sprintf("0")),
		tgbotapi.NewInlineKeyboardButtonData(emoji.RightArrow.String(), fmt.Sprintf("+12")),
	}
	wasMessageIDFound := false
	v, _ = state.RWMapInstance.GetConfig(v.UserID)
	if update.CallbackQuery != nil {
		v.MessageID = update.CallbackQuery.Message.MessageID
		wasMessageIDFound = true
	}
	state.RWMapInstance.UpdateConfig(v)

	markup := tgbotapi.NewInlineKeyboardMarkup(arrRows...)

	if wasMessageIDFound {
		msg := tgbotapi.NewEditMessageReplyMarkup(v.ChatID, update.CallbackQuery.Message.MessageID, markup)
		r.bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(v.ChatID, "Выберите год")
		msg.ReplyMarkup = markup
		if _, err := r.bot.Send(msg); err != nil {
			state.Logger.Error(err.Error(), zap.Int64("chatID", v.ChatID))
		}
	}

}

func (r *BirthdayRouter) addMonth(update tgbotapi.Update) {

	v, _ := state.RWMapInstance.GetConfig(update.SentFrom().ID)
	v.CurrentState = state.MONTH
	state.RWMapInstance.UpdateConfig(v)

	message := tgbotapi.NewMessage(v.ChatID, "Выберите месяц")
	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(state.GetMonthButtons()...)
	r.bot.Send(message)
}

func (r *BirthdayRouter) addDay(update tgbotapi.Update) {

	v, _ := state.RWMapInstance.GetConfig(update.SentFrom().ID)
	v.CurrentState = state.DAY
	state.RWMapInstance.UpdateConfig(v)
	message := tgbotapi.NewMessage(v.ChatID, "Выберите день")
	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(state.GenerateDaysButtons(v.Month, v.Year)...)
	r.bot.Send(message)
}
