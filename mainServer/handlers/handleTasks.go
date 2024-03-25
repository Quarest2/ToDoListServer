package handlers

import (
	"ToDoList/ToDoListServer/mainServer/pg_storage"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTask(c echo.Context) error {
	task := pg_storage.Task{}
	err := c.Bind(&task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newTask, err := pg_storage.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newTask)
}

func EditTask(c echo.Context) error {
	taskStr := c.Param("task")
	fmt.Println(taskStr)

	task := pg_storage.Task{}
	err := c.Bind(&task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	updatedTask, err := pg_storage.EditTask(task, taskStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c echo.Context) error {
	taskStr := c.Param("task")
	fmt.Println(taskStr)

	deletedUser, err := pg_storage.DeleteTask(taskStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, deletedUser)
}
