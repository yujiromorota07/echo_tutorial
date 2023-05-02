package dsmysql

import (
	entity "api/entities"
	"context"
)

type TodoDatasource interface {
	Select(ctx context.Context) ([]entity.Todo, error)
	Insert(ctx context.Context, e entity.Todo) (entity.Todo, error)
	Update(ctx context.Context, e entity.Todo) error
	Delete(ctx context.Context, e entity.TodoID) error
}
