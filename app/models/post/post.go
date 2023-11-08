package post

import (
	"database/sql"
	"github.com/feilongjump/api.howio.world/app/models"
	"github.com/feilongjump/api.howio.world/app/models/content"
	"github.com/feilongjump/api.howio.world/app/models/user"
)

type Post struct {
	models.BaseModel
	UserId uint64 `json:"-"`

	Title       string       `json:"title"`
	Excerpt     string       `json:"excerpt"`
	PublishedAt sql.NullTime `json:"published_at"`

	User    user.User       `gorm:"foreignKey:UserId" json:"user"`
	Content content.Content `gorm:"polymorphic:Owner" json:"content"`

	models.BaseTimeModel
	models.BaseDeleteTimeModel
}
