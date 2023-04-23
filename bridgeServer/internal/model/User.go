package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	userId      int64
	date        time.Time
	currentChat []Chat
}

func NewUser(userId int64, date time.Time, currentChat []Chat) *User {
	return &User{userId: userId, date: date, currentChat: currentChat}
}
