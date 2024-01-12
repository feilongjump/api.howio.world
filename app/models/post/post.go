package post

import (
	"database/sql"
	"github.com/feilongjump/api.howio.world/app/models"
	"github.com/feilongjump/api.howio.world/app/models/content"
	"github.com/feilongjump/api.howio.world/app/models/user"
)

type Post struct {
	models.BaseModel
	UserId uint64 `gorm:"type:bigint;not null;index" json:"-"`

	Title       string       `gorm:"column:title;type:varchar(100);not null;index" json:"title"`
	PublishedAt sql.NullTime `gorm:"column:published_at" json:"published_at"`

	User    user.User       `gorm:"foreignKey:UserId" json:"user"`
	Content content.Content `gorm:"polymorphic:Owner" json:"content"`

	models.BaseTimeModel
	models.BaseDeleteTimeModel
}
