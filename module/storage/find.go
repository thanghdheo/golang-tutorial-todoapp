package todostorage

import (
	"context"
	todomodel "todo-app/module/model"
)

func (sql *sqlStorage) FindTodo(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*todomodel.Todo, error) {
	var data todomodel.Todo

	if err := sql.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
