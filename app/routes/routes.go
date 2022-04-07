package routes

import (
	"todo/controller/todos"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	TodosController todos.TodosController
}

func (cl *ControllerList) RouteList(e *echo.Echo) {
	todo := e.Group("todo")
	todo.GET("/list/", cl.TodosController.GetTodos)
	todo.POST("/create", cl.TodosController.Create)
	todo.PUT("/:id", cl.TodosController.Update)
	todo.DELETE("/:id", cl.TodosController.Delete)
}