package inputport

import (
	entity "api/entities"
)

type Todo interface {
	GetTodos() ([]entity.Todo, error)
	CreateTodo(e entity.Todo) (entity.Todo, error)
}
