package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/xprasetio/go_simple_forum.git/internal/model/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{Engine: api, membershipSvc: membershipSvc}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/signup", h.SignUp)
	route.POST("/login", h.Login)
}
