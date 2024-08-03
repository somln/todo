package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/db"
	"todo/models"
	"todo/utils"
)

type CreateTodoInput struct {
	Content string `json:"content" binding:"required"`
}

type UpdateTodoContentInput struct {
	Content string `json:"content" binding:"required"`
}

type UpdateTodoStatusInput struct {
	Status models.TodoStatus `json:"status" binding:"required"`
}

func getTodoByID(c *gin.Context, todo *models.Todo) error {
	if err := db.Db.First(todo, c.Param("todoId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return err
	}
	return nil
}

func bindJSONAndCheckError(c *gin.Context, input interface{}) error {
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return err
	}
	return nil
}

func CreateTodo(c *gin.Context) {

	var input CreateTodoInput

	if err := bindJSONAndCheckError(c, &input); err != nil {
		return
	}

	todo := models.Todo{Content: input.Content, Status: models.Progress}
	db.Db.Create(&todo)

	c.JSON(http.StatusCreated, todo.TodoID)
}

func GetTodo(c *gin.Context) {

	var todo models.Todo

	if err := getTodoByID(c, &todo); err != nil {
		return
	}

	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	status := c.Query("status")

	if status != "" {
		todoStatus := models.TodoStatus(status)

		if !utils.ValidateTodoStatus(c, todoStatus) {
			return
		}
		db.Db.Where("status = ?", todoStatus).Find(&todos)

	} else {
		db.Db.Find(&todos)
	}

	c.JSON(http.StatusOK, todos)
}

func UpdateTodoContent(c *gin.Context) {

	var todo models.Todo

	if err := getTodoByID(c, &todo); err != nil {
		return
	}

	var input UpdateTodoContentInput

	if err := bindJSONAndCheckError(c, &input); err != nil {
		return
	}

	db.Db.Model(&todo).Update("Content", input.Content)
	c.JSON(http.StatusOK, todo)
}

func UpdateTodoStatus(c *gin.Context) {

	var todo models.Todo

	if err := getTodoByID(c, &todo); err != nil {
		return
	}

	var input UpdateTodoStatusInput

	if err := bindJSONAndCheckError(c, &input); err != nil {
		return
	}

	if !utils.ValidateTodoStatus(c, input.Status) {
		return
	}

	db.Db.Model(&todo).Update("Status", input.Status)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {

	var todo models.Todo

	if err := getTodoByID(c, &todo); err != nil {
		return
	}

	db.Db.Delete(&todo)
	c.JSON(http.StatusNoContent, gin.H{})
}
