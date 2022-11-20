package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AllTodos(c *gin.Context) {
	var todos []Todo
	DB.Find(&todos)

	c.IndentedJSON(http.StatusOK, todos)
}

func ChangeTodo(c *gin.Context) {
	id := c.Param("id")

	var status struct {
		Done bool
	}

	c.Bind(&status)

	var todo Todo
	DB.First(&todo, "id =?", id)

	if todo.ID.String() == "00000000-0000-0000-0000-000000000000" {
		log.Info("todo not found")
		c.Status(http.StatusNotFound)
		return
	}

	DB.Model(&todo).Updates(Todo{
		Done: status.Done,
	})

	c.Writer.WriteHeader(http.StatusNoContent)
}

func CreateTodo(c *gin.Context) {
	var todoIn struct {
		Text string
	}

	c.Bind(&todoIn)

	todo := Todo{Text: todoIn.Text, Done: false}

	result := DB.Create(&todo)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.Header("Location", "/api/v1/todo/"+todo.ID.String())
	c.IndentedJSON(http.StatusCreated, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	var todo Todo
	DB.First(&todo, "id =?", id)

	if todo.ID.String() == "00000000-0000-0000-0000-000000000000" {
		log.Info("todo not found")
		c.Status(http.StatusNotFound)
		return
	}

	DB.Delete(&Todo{}, "id =?", id)

	c.Writer.WriteHeader(http.StatusNoContent)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	log.Info("about to search for " + id)

	var todo Todo
	DB.First(&todo, "id =?", id)

	if todo.ID.String() == "00000000-0000-0000-0000-000000000000" {
		log.Info("todo not found")
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}
