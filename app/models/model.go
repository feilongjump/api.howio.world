package models

import (
	"database/sql/driver"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
}

type BaseTimeModel struct {
	CreatedAt time.Time `gorm:"column:created_at;index" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;index" json:"updated_at"`
}

type BaseDeleteTimeModel struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

type HumansTime struct {
	Humans   string `json:"humans,omitempty"`
	Datetime string `json:"datetime"`
}

// Scan implements the [Scanner] interface.
func (humansTime *HumansTime) Scan(value any) error {

	if value == nil {
		humansTime.Humans, humansTime.Datetime = "暂未发布", ""
		return nil
	}

	var valueCarbon carbon.Carbon

	switch v := value.(type) {
	case time.Time:
		valueCarbon = carbon.CreateFromStdTime(v)
	case []byte:
		valueCarbon = carbon.Parse(string(v))
	}
	if !valueCarbon.IsInvalid() {
		humansTime.Humans = valueCarbon.SetLocale("zh-CN").DiffForHumans()
		humansTime.Datetime = valueCarbon.ToDateTimeString()
	}

	return nil
}

// Value implements the [driver.Valuer] interface.
func (humansTime HumansTime) Value() (driver.Value, error) {

	if humansTime.Datetime == "" {
		return nil, nil
	}

	return carbon.Parse(humansTime.Datetime).StdTime(), nil
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
