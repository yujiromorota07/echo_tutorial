package repository

import (
	entity "api/entities"
	"context"
)

type TodoRepository interface {
	GetTodos(ctx context.Context) ([]entity.Todo, error)
	CreateTodo(ctx context.Context, e entity.Todo) (entity.Todo, error)
	UpdateTodo(ctx context.Context, e entity.Todo) error
}
