package todo

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {

	router.GET("/", GetTodos)
	router.GET("/:id", GetTodo)
	router.POST("/", CreateTodo)
	router.PUT("/:id", UpdateTodo)
	router.DELETE("/:id", DeleteTodo)
}