package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xprasetio/go_simple_forum.git/internal/model/memberships"
)

func (h *Handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt64("user_id")
	accessToken, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
