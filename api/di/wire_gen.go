// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"api/controllers/echoapi"
	"api/gateways/datasource/dsmysqlimpl"
	todo3 "api/gateways/repository/todo"
	"api/generated/todo"
	"api/usecase/inputport"
	todo2 "api/usecase/interactor/todo"
	"api/usecase/repository"
)

// Injectors from wire.go:

func IntializeTodoController() todo.ServerInterface {
	inputportTodo := IntializeTodoUsecase()
	serverInterface := echoapi.NewTodoController(inputportTodo)
	return serverInterface
}

func IntializeTodoUsecase() inputport.Todo {
	todoRepository := IntializeTodoRepository()
	inputportTodo := todo2.NewTodoUsecase(todoRepository)
	return inputportTodo
}

func IntializeTodoRepository() repository.TodoRepository {
	todoDatasource := dsmysqlimpl.NewTodoDatasource()
	todoRepository := todo3.NewTodoRepository(todoDatasource)
	return todoRepository
}
