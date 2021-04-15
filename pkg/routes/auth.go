package routes

import (
	"errors"
	"log"
	"time"
	"todo_gin/pkg/database"
	"todo_gin/pkg/models"
	"todo_gin/pkg/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const identityKey = "id"

func GetAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized:    unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			identityKey: v.ID,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.User{
		ID: uint(claims[identityKey].(float64)),
	}
}

func authenticator(c *gin.Context) (interface{}, error) {
	var user, requestUser models.User

	if err := c.ShouldBindJSON(&requestUser); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	db := database.GetConnectionByDB()
	result := db.First(&user, "email = ?", requestUser.Email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, jwt.ErrFailedAuthentication
	}

	if utils.DoPasswordsMatch(user.Password, requestUser.Password) {
		return &user, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func authorizator(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*models.User); ok {
		return true
	}
	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
