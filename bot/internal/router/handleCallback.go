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
	state.RWapInstance.Mutex.RLock()
	stateCFG, ok := state.RWapInstance.UserStateConfigs[callback.From.ID]
	state.RWapInstance.Mutex.RUnlock()

	if len(callback.Data) == 3 && stateCFG.CurrentState == state.YEAR {
		intRepr, _ := strconv.Atoi(callback.Data)
		stateCFG.Offset -= intRepr
		state.RWapInstance.Mutex.Lock()
		state.RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
		state.RWapInstance.Mutex.Unlock()
		r.addYear(update)
		return
	}
	if len(callback.Data) == 1 && stateCFG.CurrentState == state.YEAR {
		stateCFG.Year = YearIfOmit
		state.RWapInstance.Mutex.Lock()
		state.RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
		state.RWapInstance.Mutex.Unlock()
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
			state.RWapInstance.Mutex.Lock()
			state.RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
			state.RWapInstance.Mutex.Unlock()
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
			state.RWapInstance.Mutex.Lock()
			state.RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
			state.RWapInstance.Mutex.Unlock()
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
			state.RWapInstance.Mutex.Lock()
			state.RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
			state.RWapInstance.Mutex.Unlock()

			messageToBeDeleted := tgbotapi.NewDeleteMessage(stateCFG.ChatID, update.CallbackQuery.Message.MessageID)
			if _, err := r.bot.Request(messageToBeDeleted); err != nil {
				state.Logger.Error(err.Error(), zap.Int("messageID", stateCFG.MessageID))
			}
			successMessage := tgbotapi.NewMessage(stateCFG.ChatID, fmt.Sprintf("Ваши(%s) данные успешно записаны в БД %s", callback.From.FirstName, emoji.CheckBoxWithCheck.String()))
			r.bot.Send(successMessage)

			//TODO Add user's info to the DB

			state.RWapInstance.Mutex.Lock()
			delete(state.RWapInstance.UserStateConfigs, callback.From.ID)
			state.RWapInstance.Mutex.Unlock()
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
