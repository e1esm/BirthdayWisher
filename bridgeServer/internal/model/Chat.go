package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ID     int64 `gorm:"primaryKey"`
	UserId int64
}

func NewChat(chatId int64, UserId int64) *Chat {
	return &Chat{ID: chatId, UserId: UserId}
}
