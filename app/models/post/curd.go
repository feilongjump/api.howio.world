package post

import (
	"time"

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

func Get(id, userId uint64) (post Post, err error) {
	result := database.DB.Model(&Post{})

	if userId > 0 {
		// 已登录时，可获取自己的 post 或者已发布的
		result.Where("user_id = ? OR published_at <= ?", userId, time.Now())
	} else {
		// 未登录时，必须是已发布的 post，才可获取
		result.Where("published_at <= ?", time.Now())
	}

	err = result.Preload(clause.Associations).
		First(&post, id).
		Error

	return
}

func GetPaginate(ctx *gin.Context, userId uint64) (post []Post, total int64) {
	result := database.DB.Model(&Post{})

	if userId > 0 {
		result.Where("user_id = ?", userId)
	}

	result.Or("published_at <= ?", time.Now()).
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
