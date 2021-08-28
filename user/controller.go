package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []User
	err := GetAllUsersSer(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	err := GetUserSer(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	err := CreateUserSer(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	err := GetUserSer(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"user": "null",
		})
	}
	c.BindJSON(&user)
	err = UpdateUserSer(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := DeleteUserSer(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "deleted",
		})
	}
}