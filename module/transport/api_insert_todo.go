package todotransport

import (
	"net/http"
	todomodel "todo-app/module/model"
	todostorage "todo-app/module/storage"
	todobiz "todo-app/module/biz"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTodo(db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var data todomodel.TodoCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}

		storage := todostorage.NewSQLStore(db)
		biz := todobiz.NewCreateRestaurantBiz(storage)

		if err := biz.CreateTodo(ctx,&data); err != nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK,gin.H{"data": data.Id})

	}
}