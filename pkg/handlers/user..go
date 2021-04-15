package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"todo_gin/pkg/database"
	"todo_gin/pkg/models"
	"todo_gin/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = utils.HashPassword(user.Password)

	db := database.GetConnectionByDB()
	results := db.Create(&user)

	if results.Error != nil {
		c.JSON(200, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Registered",
	})
}

func Login(c *gin.Context) {
	var user, requestUser models.User

	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetConnectionByDB()

	result := db.First(&user, "email = ?", requestUser.Email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not Found"})
		return
	}

	if utils.DoPasswordsMatch(user.Password, requestUser.Password) {
		fmt.Println(user.Password)
		fmt.Println(requestUser.Password)

		c.JSON(200, gin.H{
			"message": "Sucessfull Login",
		})
	} else {
		c.JSON(401, gin.H{
			"message": "Error in Login",
		})
	}
}

func ListUsers(c *gin.Context) {
	var users []struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	db := database.GetConnectionByDB()

	results := db.Model(&models.User{}).Select("id, name, email").Find(&users)

	if results.Error != nil {
		c.JSON(200, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, &users)
}
