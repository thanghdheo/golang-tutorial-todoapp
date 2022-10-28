package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
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
			todo.POST("/create", createTodo(db))
			todo.GET("/:id", getTodo(db))
			todo.GET("", getTodos(db))
			todo.PUT("/update/:id", updateTodo(db))
			todo.PUT("/delete/:id", deleteTodo(db))
		}
	}
	router.Run(":3000")
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
	return "todo"
}

type TodoCreate struct {
	Id        int        `json:"id" gorm:"id;"`
	Title     string     `json:"title" gorm:"title;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

type TodoUpdate struct {
	Id        *int       `json:"id" gorm:"id;"`
	Title     *string    `json:"title" gorm:"title;"`
	Deleted   bool       `json:"deleted" gorm:"deleted;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (TodoCreate) TableName() string {
	return Todo{}.TableName()
}


func (TodoUpdate) TableName() string {
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

func getTodo(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data Todo

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data})

	}
}

func getTodos(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type Paging struct {
			Page  int   `json:"page" form:"page;"`
			Limit int   `json:"limit" form:"limit;"`
			Total int64 `json:"total" form:"-"`
		}

		var paging Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if paging.Page == 0 {
			paging.Page = 1
		}

		if paging.Limit <= 10 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var data []Todo

		if err := db.Table(Todo{}.TableName()).Count(&paging.Total).Offset(offset).Find(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data, "paging": paging})

	}
}

func updateTodo(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data TodoUpdate

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Deleted = false

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func deleteTodo(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data TodoUpdate

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Deleted = true

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func (res *TodoCreate) Validate() error {
	res.Id = 0
	res.UpdatedAt = nil
	res.Title = strings.TrimSpace(res.Title)

	if len(res.Title) == 0 {
		return errors.New("Title can be blank")
	}
	return nil
}
