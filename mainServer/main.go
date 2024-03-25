package main

import (
	"ToDoList/ToDoListServer/mainServer/handlers"
	"ToDoList/ToDoListServer/mainServer/pg_storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)

	pg_storage.Connect()

	e.Use(middleware.Logger(), middleware.Recover())

	e.POST("/tasks", handlers.CreateTask)
	e.PUT("/tasks/:task", handlers.EditTask)
	e.DELETE("/tasks/:task", handlers.DeleteTask)

	e.Logger.Fatal(e.Start(":8080"))

}
