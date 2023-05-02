package inputport

import (
	entity "api/entities"
)

type Todo interface {
	GetTodos() ([]entity.Todo, error)
	GetTodo(e entity.TodoID) (entity.Todo, error)
	CreateTodo(e entity.Todo) (entity.Todo, error)
	UpdateTodo(e entity.Todo) (entity.Todo, error)
	DeleteTodo(e entity.TodoID) error
}
