package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/db"
	"todo/models"
)

func CreateTodo(c *gin.Context) {

	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Content: input.Content, Status: "Progress"}
	db.Db.Create(&todo)

	c.JSON(http.StatusCreated, todo.TodoID)
}

func GetTodo(c *gin.Context) {

	var todo models.Todo

	if err := db.Db.First(&todo, c.Param("todoId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "해당 todo는 존재하지 않습니다. "})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {

	var todos []models.Todo

	status := c.Query("status")

	if status != "" {
		db.Db.Where(status).Find(&todos)
	} else {
		db.Db.Find(&todos)
	}

	c.JSON(http.StatusOK, todos)
}

func UpdateTodoContent(c *gin.Context) {

	var todo models.Todo

	if err := db.Db.First(&todo, c.Param("todoId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "해당 todo는 존재하지 않습니다. "})
		return
	}

	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db.Db.Model(&todo).Update("Content", input.Content)
	c.JSON(http.StatusOK, todo)
}

func UpdateTodoStatus(c *gin.Context) {

	var todo models.Todo

	if err := db.Db.First(&todo, c.Param("todoId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "해당 todo는 존재하지 않습니다. "})
		return
	}

	var input struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db.Db.Model(&todo).Update("Status", input.Status)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {

	var todo models.Todo

	if err := db.Db.First(&todo, c.Param("todoId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "해당 todo는 존재하지 않습니다. "})
		return
	}

	db.Db.Delete(&todo)
	c.JSON(http.StatusNoContent, gin.H{})

}
