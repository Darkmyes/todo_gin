package handlers

import (
	"net/http"
	"time"
	"todo_gin/pkg/database"
	"todo_gin/pkg/models"

	"github.com/gin-gonic/gin"
)

func RegisterTask(c *gin.Context) {
	var task models.TaskUsers

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetConnectionByDB()
	db.Create(&task)

	c.JSON(200, gin.H{
		"message": "Task Registered",
	})
}

func ListTasks(c *gin.Context) {
	//var tasks []models.Task
	var tasks []struct {
		ID          uint      `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	db := database.GetConnectionByDB()

	results := db.Model(&models.Task{}).Select("id, name, description, created_at, updated_at").Find(&tasks)
	//results := db.Find(&tasks)

	if results.Error != nil {
		c.JSON(200, gin.H{
			"message": "Error",
		})
	}

	c.JSON(200, &tasks)
}
