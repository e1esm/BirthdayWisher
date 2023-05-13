package router

import (
	state "BirthdayWisherBot/utils"
	"fmt"
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"strconv"
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
		r.addYear(update, &update.CallbackQuery.Message.MessageID)
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
			stateCFG.Year, _ = strconv.Atoi(callback.Data)
			state.RWapInstance.Mutex.Lock()
			state.RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
			state.RWapInstance.Mutex.Unlock()
			messageToBeDeleted := tgbotapi.NewDeleteMessage(stateCFG.ChatID, update.CallbackQuery.Message.MessageID)
			if _, err := r.bot.Request(messageToBeDeleted); err != nil {
				state.Logger.Error(err.Error(), zap.Int("messageID", stateCFG.MessageID))
			}
			r.addMonth(update)
		case state.MONTH:
			stateCFG.Month, _ = strconv.Atoi(callback.Data)
			state.RWapInstance.Mutex.Lock()
			state.RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
			state.RWapInstance.Mutex.Unlock()
			messageToBeDeleted := tgbotapi.NewDeleteMessage(stateCFG.ChatID, update.CallbackQuery.Message.MessageID)
			if _, err := r.bot.Request(messageToBeDeleted); err != nil {
				state.Logger.Error(err.Error(), zap.Int("messageID", stateCFG.MessageID))
			}
			r.addDay(update)

		case state.DAY:
			stateCFG.Day, _ = strconv.Atoi(callback.Data)
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
