package router

import (
	state "BirthdayWisherBot/utils"
	"fmt"
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var YearIfOmit = 1000

func (r *BirthdayRouter) handleCallback(update tgbotapi.Update) {
	callback := update.CallbackQuery

	stateCFG, ok := state.RWMapInstance.GetConfig(update.SentFrom().ID)

	if len(callback.Data) == 3 && stateCFG.CurrentState == state.YEAR {
		intRepr, _ := strconv.Atoi(callback.Data)
		stateCFG.Offset -= intRepr
		state.RWMapInstance.UpdateConfig(stateCFG)
		r.addYear(update)
		return
	}
	if len(callback.Data) == 1 && stateCFG.CurrentState == state.YEAR {
		stateCFG.Year = YearIfOmit
		state.RWMapInstance.UpdateConfig(stateCFG)
		r.addMonth(update)
		return
	}

	if ok {
		switch stateCFG.CurrentState {
		case state.YEAR:
			tempYear, _ := strconv.Atoi(callback.Data)
			if !isValidYear(tempYear) {
				alert := tgbotapi.NewCallbackWithAlert(update.CallbackQuery.ID, "Выбранный год больше текущего")
				r.bot.Send(alert)
				return
			}
			stateCFG.Year = tempYear
			state.RWMapInstance.UpdateConfig(stateCFG)
			messageToBeDeleted := tgbotapi.NewDeleteMessage(stateCFG.ChatID, update.CallbackQuery.Message.MessageID)
			if _, err := r.bot.Request(messageToBeDeleted); err != nil {
				state.Logger.Error(err.Error(), zap.Int("messageID", stateCFG.MessageID))
			}
			r.addMonth(update)
		case state.MONTH:
			tempMonth, _ := strconv.Atoi(callback.Data)
			if !isValidMonthAndYear(tempMonth, stateCFG.Year) {
				alert := tgbotapi.NewCallbackWithAlert(update.CallbackQuery.ID, "Выбранный месяц больше текущего")
				r.bot.Send(alert)
				return
			}
			stateCFG.Month = tempMonth
			state.RWMapInstance.UpdateConfig(stateCFG)
			messageToBeDeleted := tgbotapi.NewDeleteMessage(stateCFG.ChatID, update.CallbackQuery.Message.MessageID)
			if _, err := r.bot.Request(messageToBeDeleted); err != nil {
				state.Logger.Error(err.Error(), zap.Int("messageID", stateCFG.MessageID))
			}
			r.addDay(update)

		case state.DAY:

			tempDay, _ := strconv.Atoi(callback.Data)
			if !isValidDay(stateCFG.Year, stateCFG.Month, tempDay) {
				alert := tgbotapi.NewCallbackWithAlert(update.CallbackQuery.ID, "Выбранный месяц больше текущего")
				r.bot.Send(alert)
				return
			}

			stateCFG.Day = tempDay
			state.RWMapInstance.UpdateConfig(stateCFG)

			messageToBeDeleted := tgbotapi.NewDeleteMessage(stateCFG.ChatID, update.CallbackQuery.Message.MessageID)
			if _, err := r.bot.Request(messageToBeDeleted); err != nil {
				state.Logger.Error(err.Error(), zap.Int("messageID", stateCFG.MessageID))
			}

			err := r.ConnectorService.CreateInstanceToBeDelivered(update)

			if err != nil {
				failedMessage := tgbotapi.NewMessage(stateCFG.ChatID, fmt.Sprintf("Ваши данные не удалось сохранить %s", emoji.CrossMark))
				r.bot.Send(failedMessage)
			} else {
				successMessage := tgbotapi.NewMessage(stateCFG.ChatID, fmt.Sprintf("Ваши(%s) данные успешно сохранены %s", callback.From.FirstName, emoji.CheckBoxWithCheck.String()))
				r.bot.Send(successMessage)
			}
			state.RWMapInstance.DeleteConfig(callback.From.ID)
			return
		}

	}

}

func isValidYear(year int) bool {
	if year > time.Now().Year() {
		return false
	}
	return true
}

func isValidMonthAndYear(month, year int) bool {
	if year == time.Now().Year() && time.Now().Month() < time.Month(month) {
		return false
	}
	return true
}

func isValidDay(year, month, day int) bool {
	if year == time.Now().Year() && time.Now().Month() == time.Month(month) && day > time.Now().Day() {
		return false
	}
	return true
}
