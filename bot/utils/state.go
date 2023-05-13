package utils

import "sync"

func init() {
	RWapInstance = RWMap{UserStateConfigs: make(map[int64]StateConfig), Mutex: &sync.RWMutex{}}
}

type State int

const (
	START State = iota
	YEAR
	MONTH
	DAY
)

var RWapInstance RWMap

type RWMap struct {
	UserStateConfigs map[int64]StateConfig
	Mutex            *sync.RWMutex
}

type StateConfig struct {
	Date         string
	Year         int
	Month        int
	Day          int
	ChatID       int64
	UserID       int64
	MessageID    int
	CurrentState State
	Offset       int
}
