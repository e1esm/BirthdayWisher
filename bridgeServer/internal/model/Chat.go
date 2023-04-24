package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ChatId int64
	UserId int64
}

func NewChat(chatId int64, UserId int64) *Chat {
	return &Chat{ChatId: chatId, UserId: UserId}
}
