package repository

import (
	entity "api/entities"
	"context"
)

type TodoRepository interface {
	GetTodos(ctx context.Context) ([]entity.Todo, error)
	GetTodo(ctx context.Context, e entity.TodoID) (entity.Todo, error)
	CreateTodo(ctx context.Context, e entity.Todo) (entity.Todo, error)
	UpdateTodo(ctx context.Context, e entity.Todo) error
	DeleteTodo(ctx context.Context, e entity.TodoID) error
}
