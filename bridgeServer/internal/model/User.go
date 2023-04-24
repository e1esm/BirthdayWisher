package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserId      int64
	Date        time.Time
	CurrentChat []Chat
}

func NewUser(userId int64, date time.Time, currentChat []Chat) *User {
	return &User{UserId: userId, Date: date, CurrentChat: currentChat}
}
