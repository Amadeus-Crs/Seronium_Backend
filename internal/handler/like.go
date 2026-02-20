package handler

import (
	"Seronium/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var likeService = service.NewLikeService()

func Like(c *gin.Context) {
	userID := c.GetUint64("user_id")
	var req struct {
		TargetType string `json:"target_type" binding:"required"`
		TargetID   uint64 `json:"target_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2, "msg": err.Error()})
		return
	}

	if err := likeService.Like(userID, req.TargetType, req.TargetID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "点赞成功"})
}
