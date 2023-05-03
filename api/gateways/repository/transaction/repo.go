package transaction

import (
	"api/gateways/infra/inframysql"
	"api/usecase/repository"
	"context"
)

type TransactionManager struct{}

func NewTransactionManager() repository.TransactionManager {
	return TransactionManager{}
}

func (mng TransactionManager) Do(ctx context.Context, runner func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := inframysql.Client.Begin()
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, &inframysql.TxKey, tx)

	v, err := runner(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return v, nil
}
