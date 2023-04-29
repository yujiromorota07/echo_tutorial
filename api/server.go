package main

import (
	"api/di"
	todo "api/generated/todo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	todoController := di.IntializeTodoController()
	todoAPI := e.Group("")

	todo.RegisterHandlers(todoAPI, todoController)
	e.Logger.Fatal(e.Start(":1323"))
}
