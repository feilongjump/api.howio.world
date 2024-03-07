package post

import (
	"github.com/feilongjump/api.howio.world/app/models"
	"github.com/feilongjump/api.howio.world/internal/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (post *Post) Create() error {
	if err := database.DB.Create(post); err != nil {
		return err.Error
	}

	return nil
}

func (post *Post) Update() (int64, error) {

	result := database.DB.
		Model(&post).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(post)

	return result.RowsAffected, result.Error
}

func Get(id uint64) (post Post, err error) {
	result := database.DB.Preload(clause.Associations).First(&post, id)
	err = result.Error

	return
}

func GetPaginate(ctx *gin.Context) (post []Post, total int64) {
	database.DB.Model(&post).
		Count(&total).
		Scopes(models.Paginator(ctx)).
		Order("published_at desc").
		Order("created_at desc").
		Find(&post)

	return
}

func (post *Post) Delete() error {
	result := database.DB.Select("Content").Delete(post)

	return result.Error
}
