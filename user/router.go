package user

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {

	router.GET("/", GetUsers)
	router.GET("/:id", GetUser)
	router.POST("/", CreateUser)
	router.PUT("/:id", UpdateUser)
	router.DELETE("/:id", DeleteUser)
}
