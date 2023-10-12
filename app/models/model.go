package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type BaseModel struct {
	ID uint64 `json:"id"`
}

type BaseTimeModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseDeleteTimeModel struct {
	DeletedAt gorm.DeletedAt `json:"-"`
}

// Paginator 分页器
func Paginator(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(ctx.Query("page"))
		if page <= 1 {
			page = 1
		}

		perPage, _ := strconv.Atoi(ctx.Query("per_page"))
		if perPage <= 0 {
			perPage = 20
		}

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}
