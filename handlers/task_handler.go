package handlers

import (
	"crud/models"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var tasks = map[string]models.Task{
	"1": {ID: "1", Title: "Task 1", Content: "This is the first task"},
	"2": {ID: "2", Title: "Task 2", Content: "This is the second task"},
}

func GetTasks(c echo.Context) error {
	var taskList []models.Task

	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	return c.JSON(http.StatusOK, taskList)
}

func GetTask(c echo.Context) error {
	taskID := c.Param("id")

	task, found := tasks[taskID]

	if !found {
		return c.JSON(http.StatusNotFound, "Task not found")
	}

	return c.JSON(http.StatusOK, task)
}

func CreateTask(c echo.Context) error {
	var newTask models.Task

	if err := c.Bind(&newTask); err != nil {

		fmt.Println("Error binding request body:", err)
		return err
	}

	newTask.ID = uuid.New().String()
	tasks[newTask.ID] = newTask

	return c.JSON(http.StatusCreated, newTask)
}

func UpdateTask(c echo.Context) error {
	taskID := c.Param("id")

	_, found := tasks[taskID]

	if !found {
		return c.JSON(http.StatusNotFound, "Task not found")
	}

	var updatedTask models.Task

	if err := c.Bind(&updatedTask); err != nil {
		return err
	}

	updatedTask.ID = taskID
	tasks[taskID] = updatedTask

	return c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c echo.Context) error {
	taskID := c.Param("id")

	_, found := tasks[taskID]

	if !found {
		return c.JSON(http.StatusNotFound, "Task not found")
	}

	delete(tasks, taskID)

	return c.NoContent(http.StatusNoContent)
}
