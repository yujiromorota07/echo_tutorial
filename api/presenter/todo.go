package presenter

import (
	entity "api/entities"
	"api/generated/todo"
)

func GetTodos(e []entity.Todo) []todo.Todo {
	res := make([]todo.Todo, 0)
	for _, p := range e {
		r := todo.Todo{
			Id:      int32(p.ID),
			Title:   p.Title,
			Content: p.Content,
		}
		res = append(res, r)
	}
	return res
}

func GetTodo(e entity.Todo) todo.Todo {
	return todo.Todo{
		Id:      int32(e.ID),
		Title:   e.Title,
		Content: e.Content,
	}
}
