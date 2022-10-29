package todotransport

import (
	"net/http"
	"strconv"
	todobiz "todo-app/module/biz"
	todostorage "todo-app/module/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteTodo(db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id,err := strconv.Atoi(ctx.Param("id"))

		if  err != nil {
			ctx.JSONP(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}

		storage := todostorage.NewSQLStore(db)
		biz := todobiz.NewDeleteTodoBiz(storage)

		if err := biz.DeleteTodo(ctx.Request.Context(), id); err != nil {
			ctx.JSONP(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK,gin.H{"data": true})
	}
}