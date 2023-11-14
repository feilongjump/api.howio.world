package content

import (
	"github.com/feilongjump/api.howio.world/internal/markdown"
	"gorm.io/gorm"
)

func (content *Content) BeforeSave(tx *gorm.DB) (err error) {
	content.Body = markdown.Convert(content.Markdown)

	return
}
