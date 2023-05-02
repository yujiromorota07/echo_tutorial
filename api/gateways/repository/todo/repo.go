package todo

import (
	entity "api/entities"
	"api/gateways/repository/datasource/dsmysql"
	"api/usecase/repository"
	"context"
)

type TodoRepository struct {
	dsTodo dsmysql.TodoDatasource
}

func NewTodoRepository(dsTodo dsmysql.TodoDatasource) repository.TodoRepository {
	return &TodoRepository{dsTodo: dsTodo}
}

func (repo TodoRepository) GetTodos(ctx context.Context) ([]entity.Todo, error) {
	todos, err := repo.dsTodo.Select(ctx)
	if err != nil {
		return []entity.Todo{}, err
	}

	return todos, nil
}

func (repo TodoRepository) CreateTodo(ctx context.Context, e entity.Todo) (entity.Todo, error) {
	todo, err := repo.dsTodo.Insert(ctx, e)
	if err != nil {
		return entity.Todo{}, err
	}
	return todo, nil
}

func (repo TodoRepository) UpdateTodo(ctx context.Context, e entity.Todo) error {
	err := repo.dsTodo.Update(ctx, e)
	if err != nil {
		return err
	}

	return nil
}
