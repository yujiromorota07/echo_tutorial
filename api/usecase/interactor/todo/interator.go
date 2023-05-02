package todo

import (
	entity "api/entities"
	"api/usecase/inputport"
	"api/usecase/repository"
	"context"
)

type TodoUsecase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(todoRepository repository.TodoRepository) inputport.Todo {
	return &TodoUsecase{
		todoRepository: todoRepository,
	}
}

func (usecase TodoUsecase) GetTodos() ([]entity.Todo, error) {
	ctx := context.Background()
	todos, err := usecase.todoRepository.GetTodos(ctx)
	if err != nil {
		return []entity.Todo{}, err
	}

	return todos, nil
}

func (usecase TodoUsecase) CreateTodo(e entity.Todo) (entity.Todo, error) {
	ctx := context.Background()
	todo, err := usecase.todoRepository.CreateTodo(ctx, e)
	if err != nil {
		return entity.Todo{}, err
	}

	return todo, nil
}

func (usecase TodoUsecase) UpdateTodo(e entity.Todo) (entity.Todo, error) {
	ctx := context.Background()
	err := usecase.todoRepository.UpdateTodo(ctx, e)
	if err != nil {
		return entity.Todo{}, err
	}

	return e, nil
}

func (usecase TodoUsecase) DeleteTodo(e entity.TodoID) error {
	ctx := context.Background()
	err := usecase.todoRepository.DeleteTodo(ctx, e)
	if err != nil {
		return err
	}

	return nil
}
