package todobiz

import (
	"context"
	todomodel "todo-app/module/model"
)

type UpdateTodoStorage interface{
	FindTodo(ctx context.Context,cond map[string]interface{},moreKeys ...string) (*todomodel.Todo,error)

	UpdateTodo(ctx context.Context, cond map[string]interface{}, data *todomodel.TodoUpdate) error
}


type updateTodoBiz struct{
	store UpdateTodoStorage
}

func NewUpdateBiz(store UpdateTodoStorage) *updateTodoBiz{
	return &updateTodoBiz{store: store}
}

func (biz *updateTodoBiz) UpdateTodo(ctx context.Context, id int, data *todomodel.TodoUpdate) error{
	_,err := biz.store.FindTodo(ctx,map[string]interface{}{"id": id})

	if err != nil{
		return err
	}

	if err := biz.store.UpdateTodo(ctx,map[string]interface{}{"id": id}, data); err != nil{
		return err
	}

	return nil
}