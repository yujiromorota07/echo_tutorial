package dsmysqlimpl

import (
	entity "api/entities"
	"api/gateways/repository/datasource/dsmysql"
	"context"
)

type todoDatasource struct{}

func NewTodoDatasource() dsmysql.TodoDatasource {
	return &todoDatasource{}
}

func (ds todoDatasource) Select(ctx context.Context) ([]entity.Todo, error) {
	res := make([]entity.Todo, 0)
	res = append(res, entity.Todo{
		ID:      1,
		Title:   "title",
		Content: "content",
	})
	return res, nil
}

func (ds todoDatasource) Insert(ctx context.Context, e entity.Todo) (entity.Todo, error) {
	todo := entity.Todo{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
	return todo, nil
}
