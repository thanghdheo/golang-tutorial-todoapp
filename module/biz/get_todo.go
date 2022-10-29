package todobiz

import (
	"context"
	todomodel "todo-app/module/model"
)

type FindTodoStorage interface{
	FindTodo(ctx context.Context,cond map[string]interface{},moreKeys ...string) (*todomodel.Todo,error)
}

func NewFindTodoBiz(store FindTodoStorage) *findTodoBiz{
	return &findTodoBiz{store: store}
}

type findTodoBiz struct{
	store FindTodoStorage
}

func (biz *findTodoBiz) FindTodo(ctx context.Context,id int) (*todomodel.Todo,error){

	data, err := biz.store.FindTodo(ctx,map[string]interface{}{"id": id})

	if err != nil{
		return nil, err
	}

	return data, nil
}

