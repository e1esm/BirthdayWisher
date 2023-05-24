package utils

import "sync"

func init() {
	RWMapInstance = RWMap{UserStateConfigs: make(map[int64]StateConfig), Mutex: &sync.RWMutex{}}
}

type State int

const (
	START State = iota
	YEAR
	MONTH
	DAY
)

var RWMapInstance RWMap

type RWMap struct {
	UserStateConfigs map[int64]StateConfig
	Mutex            *sync.RWMutex
}

func (r *RWMap) DeleteConfig(UserID int64) {
	r.Mutex.Lock()
	delete(r.UserStateConfigs, UserID)
	r.Mutex.Unlock()
}

func (r *RWMap) UpdateConfig(config StateConfig) {
	r.Mutex.Lock()
	r.UserStateConfigs[config.UserID] = config
	r.Mutex.Unlock()
}

func (r *RWMap) GetConfig(UserID int64) (StateConfig, bool) {
	r.Mutex.Lock()
	v, ok := r.UserStateConfigs[UserID]
	r.Mutex.Unlock()
	return v, ok
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
