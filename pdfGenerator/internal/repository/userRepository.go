package repository

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"pdfGenerator/utils"
	"time"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{initDB()}
}

func initDB() *gorm.DB {
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_port := os.Getenv("DB_PORT")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_CONTAINER_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_password, db_name, db_port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			localization, err := time.LoadLocation("Europe/Moscow")
			if err != nil {
				utils.Logger.Fatal("Invalid time zone", zap.String("time zone", localization.String()))
			}
			return time.Now().In(localization)
		},
	})
	if err != nil {
		utils.Logger.Fatal("Can't open connection with the database", zap.String("error", err.Error()))
	}
	return db
}
