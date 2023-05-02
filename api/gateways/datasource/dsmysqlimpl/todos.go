package dsmysqlimpl

import (
	entity "api/entities"
	"api/gateways/infra/inframysql"
	"api/gateways/repository/datasource/dsmysql"
	"context"
)

type todoDatasource struct{}

func NewTodoDatasource() dsmysql.TodoDatasource {
	return &todoDatasource{}
}

const (
	querySelectTodos = "SELECT * FROM `todos`;"
	queryInsertTodo  = "INSERT INTO todos(title, content) VALUES (?,?)"
	queryUpdateTodo  = "UPDATE `todos` SET `title`=?, `content`=? WHERE `id`=?"
	queryDeleteTodo  = "DELETE FROM `todos` WHERE `id`=?"
)

func (ds todoDatasource) Select(ctx context.Context) ([]entity.Todo, error) {
	res := make([]entity.Todo, 0)

	dao := inframysql.GetDao(ctx)
	stmt, err := dao.Prepare(querySelectTodos)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r entity.Todo
		err = rows.Scan(&r.ID, &r.Title, &r.Content, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	defer rows.Close()

	return res, nil
}

func (ds todoDatasource) Insert(ctx context.Context, e entity.Todo) (entity.Todo, error) {
	dao := inframysql.GetDao(ctx)
	stmt, err := dao.Prepare(queryInsertTodo)
	if err != nil {
		return entity.Todo{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Title, e.Content)
	if err != nil {
		return entity.Todo{}, err
	}

	id, _ := result.LastInsertId()
	e.ID = uint32(id)

	return e, err
}

func (ds todoDatasource) Update(ctx context.Context, e entity.Todo) error {
	dao := inframysql.GetDao(ctx)
	stmt, err := dao.Prepare(queryUpdateTodo)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Title, e.Content, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ds todoDatasource) Delete(ctx context.Context, e entity.TodoID) error {
	dao := inframysql.GetDao(ctx)
	stmt, err := dao.Prepare(queryDeleteTodo)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e)
	if err != nil {
		return err
	}

	return nil
}
