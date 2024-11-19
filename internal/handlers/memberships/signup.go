package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xprasetio/go_simple_forum.git/internal/model/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	req := memberships.SignUpRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.membershipSvc.SignUp(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}
