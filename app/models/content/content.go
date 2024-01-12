package content

import "github.com/feilongjump/api.howio.world/app/models"

type Content struct {
	models.BaseModel

	OwnerID   int    `gorm:"type:bigint;not null;index" json:"-"`
	OwnerType string `gorm:"type:varchar(255);not null;index" json:"-"`

	// 暂时为空不返回，后续可能考虑由 markdown 转换成 html 进行存储
	Body     string `gorm:"type:longtext" json:"body,omitempty"`
	Markdown string `gorm:"type:longtext" json:"markdown"`

	models.BaseTimeModel
	models.BaseDeleteTimeModel
}
