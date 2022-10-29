package todostorage

import (
	"context"
	todomodel "todo-app/module/model"
)

func (store *sqlStorage) DeleteTodo(ctx context.Context, cond map[string]interface{}) error {
	if err := store.db.Table(todomodel.Todo{}.TableName()).Where(cond).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
