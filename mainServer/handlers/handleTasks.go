package handlers

import (
	"ToDoList/ToDoListServer/mainServer/pg_storage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
	id := c.Param("id")
	fmt.Println(id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	task := pg_storage.Task{}
	err = c.Bind(&task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	updatedTask, err := pg_storage.EditTask(task, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	deletedUser, err := pg_storage.DeleteTask(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, deletedUser)
}
