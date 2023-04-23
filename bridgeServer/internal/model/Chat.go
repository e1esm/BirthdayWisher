package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	chatId int64
	userId int64
}

func NewChat(chatId int64, userId int64) *Chat {
	return &Chat{chatId: chatId, userId: userId}
}
