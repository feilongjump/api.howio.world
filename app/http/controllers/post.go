package controllers

import (
	"github.com/feilongjump/api.howio.world/app/http/requests"
	postModel "github.com/feilongjump/api.howio.world/app/models/post"
	"github.com/feilongjump/api.howio.world/internal/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (*PostController) Index(ctx *gin.Context) {

	posts, total := postModel.GetPaginate(ctx)

	response.Success(ctx, gin.H{
		"data": posts,
		"meta": gin.H{
			"total": total,
		},
	})
}

func (postController *PostController) Show(ctx *gin.Context) {

	post, ok := postController.GetPost(ctx)
	if !ok {
		return
	}

	response.Success(ctx, post)
}

func (*PostController) Store(ctx *gin.Context) {

	params := requests.PostStoreRequest{}
	if ok := requests.Validator(ctx, &params, params.ErrorMessage()); !ok {
		return
	}

	post := postModel.Post{
		Title:  params.Title,
		UserId: ctx.MustGet("user_id").(uint64),
	}
	if err := post.Create(); err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	response.SuccessCreated(ctx, post)
}

func (postController *PostController) Update(ctx *gin.Context) {

	params := requests.PostUpdateRequest{}
	if ok := requests.Validator(ctx, &params, params.ErrorMessage()); !ok {
		return
	}

	post, ok := postController.GetPost(ctx)
	if !ok {
		return
	}

	post.Title = params.Title
	if _, err := post.Update(); err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	response.Success(ctx, post)
}

func (postController *PostController) Destroy(ctx *gin.Context) {

	post, ok := postController.GetPost(ctx)
	if !ok {
		return
	}

	if err := post.Delete(); err != nil {
		response.InternalServerError(ctx)
		return
	}

	response.SuccessNoContent(ctx)
}

// GetPost 获取 Post 数据
func (*PostController) GetPost(ctx *gin.Context) (postModel.Post, bool) {
	id, _ := strconv.Atoi(ctx.Param("post"))
	post, err := postModel.Get(uint64(id))
	if err != nil {
		response.NotFound(ctx)
		return post, false
	}

	return post, true
}
