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

