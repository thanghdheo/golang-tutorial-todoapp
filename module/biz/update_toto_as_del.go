package todobiz

import (
	"context"
	todomodel "todo-app/module/model"
)

type UpdateTodoAsDelStorage interface {
	FindTodo(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*todomodel.Todo, error)

	UpdateTodoAsDelete(ctx context.Context, cond map[string]interface{}, data *todomodel.TodoUpdate) error
}

type updateTodoAsDelBiz struct {
	store UpdateTodoAsDelStorage
}

func NewUpdateTodoAsDelBiz(store UpdateTodoAsDelStorage) *updateTodoAsDelBiz {
	return &updateTodoAsDelBiz{store: store}
}

func (biz *updateTodoAsDelBiz) UpdateTodoAsDelete(ctx context.Context, id int, data *todomodel.TodoUpdate) error {
	_, err := biz.store.FindTodo(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.store.UpdateTodoAsDelete(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil

}
