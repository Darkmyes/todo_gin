package handlers

import (
	"net/http"
	"todo_gin/configs"
	"todo_gin/pkg/database"
	"todo_gin/pkg/models"
	"todo_gin/pkg/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
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

func UserByToken(c *gin.Context) {
	var user models.User

	claims := jwt.ExtractClaims(c)
	userID := uint(claims[configs.IdentityKey].(float64))

	db := database.GetConnectionByDB()
	results := db.Model(&models.User{}).Select("id, name, email").Find(&user, userID)

	if results.Error != nil {
		c.JSON(200, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, &user)
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
