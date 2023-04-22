package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	chatId int64
	userId int64
}
