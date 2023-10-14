package requests

type PostStoreRequest struct {
	Title string `json:"title" binding:"required,gte=2,lte=64"`
}

func (PostStoreRequest) ErrorMessage() map[string]string {
	return map[string]string{
		"Title.required": "请输入标题",
		"Title.gte":      "标题长度不得小于 2",
		"Title.lte":      "标题长度不得大于 64",
	}
}

type PostUpdateRequest struct {
	Title string `json:"title" binding:"required,gte=2,lte=64"`
}

func (PostUpdateRequest) ErrorMessage() map[string]string {
	return map[string]string{
		"Title.required": "请输入标题",
		"Title.gte":      "标题长度不得小于 2",
		"Title.lte":      "标题长度不得大于 64",
	}
}
