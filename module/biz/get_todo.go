package todobiz

import "context"

type GetTodo interface{
	GetTodo(ctx context.Context,id int,)
}