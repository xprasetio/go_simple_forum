package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xprasetio/go_simple_forum.git/internal/model/memberships"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	req := memberships.LoginRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, refreshToken, err := h.membershipSvc.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := memberships.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, response)
}
