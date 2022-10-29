package todostorage

import (
	"context"
	"todo-app/common"
	todomodel "todo-app/module/model"
)


func (store *sqlStorage) GetTodos(ctx context.Context,filter *todomodel.Filter, pading *common.Paging) ([]todomodel.Todo, error){

	offset := (pading.Page - 1) * pading.Limit

	if v := filter.OwnerId; v >0 {
		store.db.Where("ownerid = ?",v)	
	}

	var data []todomodel.Todo

	if  err := store.db.Table(todomodel.Todo{}.TableName()).Count(&pading.Total).Offset(offset).Limit(pading.Limit).Find(&data).Error; err != nil {
			return nil, err
	}

	return data, nil
}