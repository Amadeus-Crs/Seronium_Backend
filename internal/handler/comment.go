package handler

import (
	"Seronium/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var commentService = service.NewCommentService()

func CreateComment(c *gin.Context) {
	userID := c.GetUint64("user_id")
	var req struct {
		TargetType string `json:"target_type" binding:"required"`
		TargetID   uint64 `json:"target_id" binding:"required"`
		Content    string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2, "msg": err.Error()})
		return
	}

	if err := commentService.Create(userID, req.TargetType, req.TargetID, req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": "评论失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "评论成功"})
}
