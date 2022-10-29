package todotransport

import (
	"net/http"
	"strconv"
	todobiz "todo-app/module/biz"
	todostorage "todo-app/module/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindTodo(db *gorm.DB)  gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return 
		}

		storage := todostorage.NewSQLStore(db)
		biz := todobiz.NewFindTodoBiz(storage)

		data, err := biz.FindTodo(ctx.Request.Context(),id)

		if err != nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return 
		}

		ctx.JSON(http.StatusOK,gin.H{"data": data})
		
	}
}