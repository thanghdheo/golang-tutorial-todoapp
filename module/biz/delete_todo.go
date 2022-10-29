package todobiz

import (
	"context"
	todomodel "todo-app/module/model"
)

type DeletetTodoStorage interface {
	FindTodo(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*todomodel.Todo, error)

	DeleteTodo(ctx context.Context, cond map[string]interface{}) error
}

func NewDeleteTodoBiz(store DeletetTodoStorage) *deleteTodoBiz {
	return &deleteTodoBiz{store: store}
}

type deleteTodoBiz struct {
	store DeletetTodoStorage
}

func (biz *deleteTodoBiz) DeleteTodo(ctx context.Context, id int) error {
	_, err := biz.store.FindTodo(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.store.DeleteTodo(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
