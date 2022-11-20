package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var start = time.Now()

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

	log.WithFields(log.Fields{
		"time_sec": time.Since(start).Seconds(),
	}).Info("Started Application")

	router.Run()
}
