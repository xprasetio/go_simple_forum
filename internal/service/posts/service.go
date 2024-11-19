package posts

import (
	"context"

	"github.com/xprasetio/go_simple_forum.git/internal/configs"
	"github.com/xprasetio/go_simple_forum.git/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
}

type service struct {
	postRepo postRepository
	cfg      *configs.Config
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{cfg: cfg, postRepo: postRepo}
}
