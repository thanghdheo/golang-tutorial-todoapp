package todobiz

import (
	"context"
	todomodel "todo-app/module/model"
)

type CreateTodoStorage interface {
	CreateTodo(ctx context.Context, data *todomodel.TodoCreate) error
}

type createBiz struct {
	store CreateTodoStorage
}

func NewCreateRestaurantBiz(store CreateTodoStorage) *createBiz{
	return &createBiz{store: store}
}

func (biz *createBiz) CreateTodo(ctx context.Context, data *todomodel.TodoCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateTodo(ctx, data); err != nil {
		return err
	}

	return nil
}
