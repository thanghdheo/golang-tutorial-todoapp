package todobiz

import (
	"context"
	"todo-app/common"
	todomodel "todo-app/module/model"
)

type GetTodosStorage interface {
	GetTodos(ctx context.Context, filter *todomodel.Filter, pading *common.Paging) ([]todomodel.Todo, error)
}

type getTodosBiz struct {
	store GetTodosStorage
}

func NewGetTodosBiz(store GetTodosStorage) *getTodosBiz {
	return &getTodosBiz{store: store}
}

func (biz *getTodosBiz) GetTodos(ctx context.Context, pading *common.Paging, filter *todomodel.Filter) ([]todomodel.Todo, error) {
	data, err := biz.store.GetTodos(ctx, filter, pading)
	if err != nil {
		return nil, err
	}

	return data, nil
}
