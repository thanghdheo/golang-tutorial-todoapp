package main

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:ead8686ba57479778a76e@tcp(127.0.0.1:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected:", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		todo := v1.Group("/todo")
		{
			todo.POST("/create",createTodo(db))
		}
	}
}

type Todo struct {
	Id        int        `json:"id" gorm:"id;"`
	Title     string     `json:"title" gorm:"title;"`
	Status    string     `json:"status" gorm:"status;"`
	Deleted   bool       `json:"deleted" gorm:"deleted"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Todo) TableName() string {
	return "todogr"
}

type TodoCreate struct {
	Id    int    `json:"id" gorm:"id;"`
	Title string `json:"title" gorm:"title;"`
}

func (TodoCreate) TableName() string {
	return Todo{}.TableName()
}

func createTodo(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data TodoCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		if err := data.Validate(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&data).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data :": data.Id})
	}
}

func (res *TodoCreate) Validate() error {
	res.Id = 0
	res.Title = strings.TrimSpace(res.Title)

	if len(res.Title) == 0 {
		return errors.New("Title can be blank")
	}
	return nil
}
