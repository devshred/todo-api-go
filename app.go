package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	LoadEnvVariables()
	ConnectToDB()
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	todoRoutes := router.Group("/api/v1/todo")
	{
		todoRoutes.GET("/", AllTodos)
		todoRoutes.PATCH("/:id", ChangeTodo)
		todoRoutes.POST("/", CreateTodo)
		todoRoutes.DELETE("/:id", DeleteTodo)
		todoRoutes.GET("/:id", GetTodo)
	}

	router.Run()
}
