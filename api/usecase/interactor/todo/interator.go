package todo

import (
	entity "api/entities"
	"api/usecase/inputport"
	"api/usecase/repository"
	"context"
)

type TodoUsecase struct {
	transaction    repository.TransactionManager
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(transaction repository.TransactionManager, todoRepository repository.TodoRepository) inputport.Todo {
	return &TodoUsecase{
		transaction:    transaction,
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

func (usecase TodoUsecase) GetTodo(e entity.TodoID) (entity.Todo, error) {
	ctx := context.Background()
	todo, err := usecase.todoRepository.GetTodo(ctx, e)
	if err != nil {
		return entity.Todo{}, err
	}

	return todo, nil
}

func (usecase TodoUsecase) CreateTodo(e entity.Todo) (entity.Todo, error) {
	ctx := context.Background()
	v, err := usecase.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		return usecase.todoRepository.CreateTodo(ctx, e)
	})

	if err != nil {
		return entity.Todo{}, err
	}

	todo := v.(entity.Todo)
	return todo, nil
}

func (usecase TodoUsecase) UpdateTodo(e entity.Todo) (entity.Todo, error) {
	ctx := context.Background()
	_, err := usecase.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		_, err := usecase.todoRepository.GetTodo(ctx, entity.TodoID(e.ID))
		if err != nil {
			return entity.Todo{}, err
		}

		return entity.Todo{}, usecase.todoRepository.UpdateTodo(ctx, e)
	})

	return e, err
}

func (usecase TodoUsecase) DeleteTodo(e entity.TodoID) error {
	ctx := context.Background()
	_, err := usecase.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		return entity.Todo{}, usecase.todoRepository.DeleteTodo(ctx, e)
	})

	return err
}
