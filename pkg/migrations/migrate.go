package migrations

import (
	"todo_gin/pkg/database"
	"todo_gin/pkg/models"
)

func Migrate() {
	db := database.GetConnectionByDB()
	db.AutoMigrate(
		&models.User{},
		&models.Task{},
		&models.UserTask{},
	)
}
