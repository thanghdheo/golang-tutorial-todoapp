package todotransport

import (
	"net/http"
	"todo-app/common"
	todobiz "todo-app/module/biz"
	todomodel "todo-app/module/model"
	todostorage "todo-app/module/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTodos(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_ = paging.Validate()

		var filter todomodel.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := todostorage.NewSQLStore(db)
		biz := todobiz.NewGetTodosBiz(storage)

		data, err := biz.GetTodos(ctx.Request.Context(),&paging,&filter)

		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data,"paging": paging})
			
	}
}
