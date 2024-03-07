package post

import (
	"github.com/feilongjump/api.howio.world/app/models"
	"github.com/feilongjump/api.howio.world/app/models/content"
	"github.com/feilongjump/api.howio.world/app/models/user"
	"github.com/golang-module/carbon/v2"
)

type Post struct {
	models.BaseModel
	UserId uint64 `gorm:"type:bigint;not null;index" json:"-"`

	Title       string            `gorm:"column:title;type:varchar(100);not null;index" json:"title"`
	PublishedAt models.HumansTime `gorm:"column:published_at" json:"published_at"`

	User    *user.User       `gorm:"foreignKey:UserId" json:"user,omitempty"`
	Content *content.Content `gorm:"polymorphic:Owner" json:"content,omitempty"`

	models.BaseTimeModel
	models.BaseDeleteTimeModel
}

// GetPublishedAt 获取发布时间
func GetPublishedAt(post *Post, publishedStr string) {

	postPublishedAt := carbon.Parse(post.PublishedAt.Datetime)
	paramsPublishedAt := carbon.Parse(publishedStr)

	if postPublishedAt.IsInvalid() && paramsPublishedAt.IsFuture() {
		// 源时间无效 && 参数时间为未来，即可赋值
		post.PublishedAt.Datetime = publishedStr
	} else if postPublishedAt.IsFuture() &&
		(paramsPublishedAt.IsInvalid() || paramsPublishedAt.IsFuture()) {
		// 源时间 && 参数时间都为未来，即可赋值
		post.PublishedAt.Datetime = publishedStr
	}

	if post.PublishedAt.Datetime != "" {
		post.PublishedAt.Humans = carbon.
			Parse(post.PublishedAt.Datetime).
			SetLocale("zh-CN").
			DiffForHumans()
	} else {
		post.PublishedAt.Humans = ""
	}

	return
}
