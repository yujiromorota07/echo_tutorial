package main

import (
	"api/di"
	todo "api/generated/todo"
	sql "database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	todoController := di.IntializeTodoController()
	todoAPI := e.Group("")
	todo.RegisterHandlers(todoAPI, todoController)

	db, err := sql.Open("mysql", "echo:echo@tcp(127.0.0.1:3306)/echodev")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()

	e.Logger.Fatal(e.Start(":1323"))
}
