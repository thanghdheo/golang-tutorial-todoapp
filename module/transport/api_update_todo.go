package todotransport

import (
	"net/http"
	"strconv"
	todobiz "todo-app/module/biz"
	todomodel "todo-app/module/model"
	todostorage "todo-app/module/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateTodo(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data todomodel.TodoUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := todostorage.NewSQLStore(db)
		biz := todobiz.NewUpdateBiz(storage)

		if err := biz.UpdateTodo(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
