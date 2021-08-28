package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/junongch/to-do-gin-gonic/config"
	"github.com/junongch/to-do-gin-gonic/user"
	"github.com/junongch/to-do-gin-gonic/todo"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitializeDB()

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		todos := v1.Group("todos")
		todo.Router(todos)

		users := v1.Group("users")
		user.Router(users)
	}

	r.Run()
}
