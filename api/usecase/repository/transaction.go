package repository

import "context"

type Transaction interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Transactions interface {
	Get(key string, builder func() (Transaction, error)) (Transaction, error)
	Succeeded(f func() error)
}

type TransactionManager interface {
	Do(ctx context.Context, runner func(ctx context.Context) (interface{}, error)) (interface{}, error)
}
