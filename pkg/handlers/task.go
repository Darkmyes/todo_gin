package handlers

import (
	"errors"
	"net/http"
	"time"
	"todo_gin/configs"
	"todo_gin/pkg/database"
	"todo_gin/pkg/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims := jwt.ExtractClaims(c)
	task.UserID = uint(claims[configs.IdentityKey].(float64))

	db := database.GetConnectionByDB()
	db.Create(&task)

	c.JSON(200, gin.H{
		"message": "Task Registered",
	})
}

func UpdateTask(c *gin.Context) {
	var task map[string]interface{}
	var originalTask models.Task

	taskID := c.Param("id")

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetConnectionByDB()
	result := db.First(&originalTask, taskID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not Found"})
		return
	}

	for field, value := range task {
		switch field {
		case "name":
			originalTask.Name = value.(string)
		case "description":
			originalTask.Description = value.(string)
		case "state":
			originalTask.State = value.(bool)
		case "user":
			originalTask.UserID = value.(uint)
		}
	}

	db.Save(&originalTask)

	c.JSON(200, gin.H{
		"message": "Task Updated",
	})
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	taskID := c.Param("id")

	db := database.GetConnectionByDB()
	result := db.First(&task, taskID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not Found"})
		return
	}

	db.Delete(&task)

	c.JSON(200, gin.H{
		"message": "Task Deleted",
	})
}

func ListTasks(c *gin.Context) {
	//var tasks []models.Task
	var tasks []struct {
		ID          uint      `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		State       bool      `json:"state"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	claims := jwt.ExtractClaims(c)
	userID := uint(claims[configs.IdentityKey].(float64))

	db := database.GetConnectionByDB()

	results := db.Model(&models.Task{}).
		Select("id, name, description, state, created_at, updated_at").
		Where("user_id = ?", userID).
		Find(&tasks)
	//results := db.Find(&tasks)

	if results.Error != nil {
		c.JSON(200, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, &tasks)
}
