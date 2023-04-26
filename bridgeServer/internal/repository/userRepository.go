package repository

import (
	"bridgeServer/internal/model"
	"errors"
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
	err := r.db.First(&retrievedUser, user.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.db.Create(user)
	} else {
		r.db.Save(user)
	}
}

func (r *UserRepository) FindUsers() []model.User {
	users := make([]model.User, 0, 10)
	r.db.Raw("select * from users where extract(month from date) = extract(month from current_date) and extract(day from date) = extract(day from current_date)").Scan(&users)
	return users
}
