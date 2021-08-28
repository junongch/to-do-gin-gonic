package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []Todo
	err := GetAllTodosSer(&todos)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func GetTodo(c *gin.Context) {
	var todo Todo
	id := c.Param("id")
	err := GetTodoSer(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	err := CreateTodoSer(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	var todo Todo
	id := c.Param("id")
	err := GetTodoSer(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"todo": "null",
		})
	}
	c.BindJSON(&todo)
	err = UpdateTodoSer(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := DeleteTodoSer(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "deleted",
		})
	}
}
