package memberships

import (
	"context"

	"github.com/xprasetio/go_simple_forum.git/internal/configs"
	"github.com/xprasetio/go_simple_forum.git/internal/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	membershipRepo membershipRepository
	cfg            *configs.Config
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{cfg: cfg, membershipRepo: membershipRepo}
}
