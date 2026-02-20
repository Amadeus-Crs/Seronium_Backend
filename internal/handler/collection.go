package handler

import (
	"Seronium/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var collectionService = service.NewCollectionService()

func Collect(c *gin.Context) {
	userID := c.GetUint64("user_id")
	var req struct {
		PostID uint64 `json:"post_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2, "msg": err.Error()})
		return
	}

	if err := collectionService.Collect(userID, req.PostID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "收藏成功"})
}
