package content

import "github.com/feilongjump/api.howio.world/app/models"

type Content struct {
	models.BaseModel

	OwnerID   int    `json:"-"`
	OwnerType string `json:"-"`

	// 暂时为空不返回，后续可能考虑由 markdown 转换成 html 进行存储
	Body     string `json:"body,omitempty"`
	Markdown string `json:"markdown"`

	models.BaseTimeModel
	models.BaseDeleteTimeModel
}
