package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/xprasetio/go_simple_forum.git/internal/middleware"
	"github.com/xprasetio/go_simple_forum.git/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error
}

type Handler struct {
	*gin.Engine
	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{Engine: api, postSvc: postSvc}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
	route.POST("/comment/:postID", h.CreateComment)
	route.PUT("/user-activity/:postID", h.UpsertUserActivity)
}
