package model

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey"`
	Date        time.Time
	Username    string
	CurrentChat []Chat `gorm:"foreignKey:UserId"`
}

func NewUser(userId int64, date time.Time, currentChat []Chat, username string) *User {
	log.Println(username + " Username of the incoming user")
	return &User{ID: userId, Date: date, CurrentChat: currentChat, Username: username}
}
