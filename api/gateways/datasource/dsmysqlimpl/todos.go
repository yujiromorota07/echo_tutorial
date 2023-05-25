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
	querySelectTodos    = "SELECT * FROM `todos`"
	querySelectTodoByID = "SELECT * FROM `todos` WHERE `id`=?"
	queryInsertTodo     = "INSERT INTO todos(title, content, status_code) VALUES (?,?,1)"
	queryUpdateTodo     = "UPDATE `todos` SET `title`=?, `content`=? ,`status_code`=? WHERE `id`=?"
	queryDeleteTodo     = "DELETE FROM `todos` WHERE `id`=?"
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
		err = rows.Scan(&r.ID, &r.Title, &r.Content, &r.StatusCode, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	defer rows.Close()

	return res, nil
}

func (ds todoDatasource) SelectById(ctx context.Context, e entity.TodoID) (entity.Todo, error) {
	var todo entity.Todo
	dao := inframysql.GetDao(ctx)
	stmt, err := dao.Prepare(querySelectTodoByID)
	if err != nil {
		return entity.Todo{}, err
	}
	defer stmt.Close()

	result := stmt.QueryRow(e)
	if getErr := result.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.StatusCode, &todo.CreatedAt, &todo.UpdatedAt); getErr != nil {
		return todo, getErr
	}

	return todo, nil
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
	e.StatusCode = 1

	return e, err
}

func (ds todoDatasource) Update(ctx context.Context, e entity.Todo) error {
	dao := inframysql.GetDao(ctx)
	stmt, err := dao.Prepare(queryUpdateTodo)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Title, e.Content, e.StatusCode, e.ID)
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
