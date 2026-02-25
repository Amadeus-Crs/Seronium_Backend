package handler

import (
	"Seronium/internal/middleware"
	"Seronium/internal/service"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

var commentService = service.NewCommentService()

func CreateComment(ctx context.Context, c *app.RequestContext) {
	userID := middleware.GetUserID(ctx, c)
	var req struct {
		TargetType string `json:"target_type" binding:"required"`
		TargetID   uint64 `json:"target_id" binding:"required"`
		Content    string `json:"content" binding:"required"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"code": 2, "msg": err.Error()})
		return
	}

	if err := commentService.Create(userID, req.TargetType, req.TargetID, req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": "评论失败"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "msg": "评论成功"})
}
