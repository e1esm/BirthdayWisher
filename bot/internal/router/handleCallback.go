package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

var YearIfOmit = "1000"

func (r *BirthdayRouter) handleCallback(update tgbotapi.Update) {
	callback := update.CallbackQuery
	RWapInstance.Mutex.RLock()
	stateCFG, ok := RWapInstance.UserStateConfigs[callback.From.ID]
	RWapInstance.Mutex.RUnlock()

	if len(callback.Data) == 3 && stateCFG.CurrentState == YEAR {
		intRepr, _ := strconv.Atoi(callback.Data)
		stateCFG.Offset -= intRepr
		RWapInstance.Mutex.Lock()
		RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
		RWapInstance.Mutex.Unlock()
		r.addYear(update, &update.CallbackQuery.Message.MessageID)
		return
	}
	if len(callback.Data) == 1 && stateCFG.CurrentState == YEAR {
		stateCFG.Date += YearIfOmit
		RWapInstance.Mutex.Lock()
		RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
		RWapInstance.Mutex.Unlock()
		//addMonth...
	}

	if ok {
		switch stateCFG.CurrentState {
		case YEAR:
			stateCFG.Date += callback.Data
			RWapInstance.Mutex.Lock()
			RWapInstance.UserStateConfigs[callback.From.ID] = stateCFG
			RWapInstance.Mutex.Unlock()
			msg := tgbotapi.NewMessage(update.FromChat().ID, callback.Data)

			r.bot.Send(msg)

		}
	}
}
