package echoapi

import (
	entity "api/entities"
	"api/generated/todo"
	"api/presenter"
	"api/usecase/inputport"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	todoUsecase inputport.Todo
}

func NewTodoController(todoUsecase inputport.Todo) todo.ServerInterface {
	return &TodoController{
		todoUsecase: todoUsecase,
	}
}

func (ctrl TodoController) GetTodos(ctx echo.Context) error {
	res, err := ctrl.todoUsecase.GetTodos()
	if err != nil {
		return presenter.ErrorResponse(ctx, http.StatusBadRequest, "todoの取得に失敗しました。")
	}

	return ctx.JSON(http.StatusOK, presenter.GetTodos(res))
}

func (ctrl TodoController) PostTodo(ctx echo.Context) error {
	req := new(todo.CreateTodo)
	err := ctx.Bind(req)
	if err != nil {
		return presenter.ErrorResponse(ctx, http.StatusBadRequest, "todoの追加に失敗しました。")
	}

	res, err := ctrl.todoUsecase.CreateTodo(entity.Todo{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return presenter.ErrorResponse(ctx, http.StatusBadRequest, "todoの追加に失敗しました。")
	}

	return ctx.JSON(http.StatusOK, presenter.GetTodo(res))
}

func (ctrl TodoController) PutTodo(ctx echo.Context) error {
	req := new(todo.Todo)
	err := ctx.Bind(req)
	if err != nil {
		return presenter.ErrorResponse(ctx, http.StatusBadRequest, "todoの編集に失敗しました。")
	}

	res, err := ctrl.todoUsecase.UpdateTodo(entity.Todo{
		ID:      uint32(req.Id),
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return presenter.ErrorResponse(ctx, http.StatusBadRequest, "todoの編集に失敗しました。")
	}

	return ctx.JSON(http.StatusOK, res)
}

func (ctrl TodoController) GetTodosId(ctx echo.Context, id int32) error {
	return nil
}

func (ctrl TodoController) DeleteTodosId(ctx echo.Context, id int32) error {
	err := ctrl.todoUsecase.DeleteTodo(entity.TodoID(id))
	if err != nil {
		return presenter.ErrorResponse(ctx, http.StatusBadRequest, "todoの削除に失敗しました。")
	}

	return presenter.SuccessResponse(ctx, "todoの削除に成功しました。")
}
