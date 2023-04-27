package dsmysql

import (
	entity "api/entities"
	"context"
)

type TodoDatasource interface {
	Select(ctx context.Context) ([]entity.Todo, error)
	Insert(ctx context.Context, e entity.Todo) (entity.Todo, error)
}
