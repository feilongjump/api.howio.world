package post

import (
	"github.com/feilongjump/api.howio.world/app/models"
	"github.com/feilongjump/api.howio.world/internal/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func (post *Post) Create() error {
	if err := database.DB.Create(&post); err != nil {
		return err.Error
	}

	return nil
}

func Get(id uint64) (post Post, err error) {
	result := database.DB.Preload(clause.Associations).First(&post, id)
	err = result.Error

	return
}

func GetPaginate(ctx *gin.Context) (post []Post, total int64) {
	database.DB.Scopes(models.Paginator(ctx)).Find(&post).Count(&total)

	return
}
