package repository

import (
	"bridgeServer/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SaveUser(user *model.User) {
	var retrievedUser model.User
	r.db.First(&retrievedUser)
	if retrievedUser.ID == 0 {
		r.db.Create(user)

	} else {
		r.db.Model(user).Association("CurrentChat").Append(user.CurrentChat)
	}
}
