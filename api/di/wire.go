//go:build wireinject
// +build wireinject

package di

import (
	echoAPICtrl "api/controllers/echoapi"
	dsmysqlimpl "api/gateways/datasource/dsmysqlimpl"
	todoRepository "api/gateways/repository/todo"
	transaction "api/gateways/repository/transaction"
	todo "api/generated/todo"
	inputport "api/usecase/inputport"
	todoUsecase "api/usecase/interactor/todo"
	repository "api/usecase/repository"

	"github.com/google/wire"
)

func IntializeTodoController() todo.ServerInterface {
	wire.Build(
		echoAPICtrl.NewTodoController,
		IntializeTodoUsecase,
	)
	return echoAPICtrl.TodoController{}
}

func IntializeTodoUsecase() inputport.Todo {
	wire.Build(
		todoUsecase.NewTodoUsecase,
		InitializeTransactionRepository,
		IntializeTodoRepository,
	)
	return todoUsecase.TodoUsecase{}
}

func IntializeTodoRepository() repository.TodoRepository {
	wire.Build(
		todoRepository.NewTodoRepository,
		dsmysqlimpl.NewTodoDatasource,
	)
	return todoRepository.TodoRepository{}
}

func InitializeTransactionRepository() repository.TransactionManager {
	wire.Build(
		transaction.NewTransactionManager,
	)
	return transaction.TransactionManager{}
}
