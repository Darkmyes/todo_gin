package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"
	"todo_gin/pkg/database"
	"todo_gin/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		case "users":
			var originalUsers []models.UserTask
			result := db.Find(&originalUsers, "task_id = ?", taskID)
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error while trying updating task"})
				return
			}

			userTasksMustDelete := []uint{}

			for _, originalUser := range originalUsers {
				exists := false

				for _, user := range value.([]uint) {
					if user == originalUser.UserID {
						exists = true
					}
				}

				if !exists {
					userTasksMustDelete = append(userTasksMustDelete, originalUser.Task.ID)
				}
			}

			result = db.Delete(&models.UserTask{}, userTasksMustDelete)

			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error while trying updating task"})
				return
			}

			for _, user := range value.([]uint) {
				exists := false

				for _, originalUser := range originalUsers {
					if user == originalUser.UserID {
						exists = true
					}
				}

				if !exists {
					taskIDUint, err := strconv.Atoi(taskID)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Error while trying updating task"})
						return
					}
					newUserTask := models.UserTask{UserID: user, TaskID: uint(taskIDUint)}
					db.Create(&newUserTask)
				}
			}

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

	db := database.GetConnectionByDB()

	results := db.Model(&models.Task{}).Select("id, name, description, state, created_at, updated_at").Find(&tasks)
	//results := db.Find(&tasks)

	if results.Error != nil {
		c.JSON(200, gin.H{
			"message": "Error",
		})
	}

	c.JSON(200, &tasks)
}
