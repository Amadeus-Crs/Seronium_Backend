package handler

import (
	"Seronium/internal/middleware"
	"Seronium/internal/service"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

var likeService = service.NewLikeService()

func Like(ctx context.Context, c *app.RequestContext) {
	userID := middleware.GetUserID(ctx, c)
	var req struct {
		TargetType string `json:"target_type" binding:"required"`
		TargetID   uint64 `json:"target_id" binding:"required"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"code": 2, "msg": err.Error()})
		return
	}

	if err := likeService.Like(userID, req.TargetType, req.TargetID); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "msg": "点赞成功"})
}
