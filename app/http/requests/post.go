package requests

type Content struct {
	Body     string `json:"body"`
	Markdown string `json:"markdown" binding:"required"`
}

type PostStoreRequest struct {
	Title   string `json:"title" binding:"required,gte=2,lte=64"`
	Content Content
}

func (PostStoreRequest) ErrorMessage() map[string]string {
	return map[string]string{
		"Title.required":    "请输入标题",
		"Title.gte":         "标题长度不得小于 2",
		"Title.lte":         "标题长度不得大于 64",
		"Markdown.required": "请输入内容",
	}
}

type PostUpdateRequest struct {
	Title   string `json:"title" binding:"required,gte=2,lte=64"`
	Content Content
}

func (PostUpdateRequest) ErrorMessage() map[string]string {
	return map[string]string{
		"Title.required":    "请输入标题",
		"Title.gte":         "标题长度不得小于 2",
		"Title.lte":         "标题长度不得大于 64",
		"Markdown.required": "请输入内容",
	}
}
