package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"re-pet/internal/models"
)

var DB *gorm.DB

func InitDB() {
	// Настройки подключения к PostgreSQL
	dsn := "host=localhost user=your_user dbname=todo_db password=your_password sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Автомиграция (создание таблицы Task)
	DB.AutoMigrate(&models.Task{})
}
