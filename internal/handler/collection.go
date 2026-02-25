package handler

import (
	"Seronium/internal/middleware"
	"Seronium/internal/service"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

var collectionService = service.NewCollectionService()

func Collect(ctx context.Context, c *app.RequestContext) {
	userID := middleware.GetUserID(ctx, c)
	var req struct {
		PostID uint64 `json:"post_id" binding:"required"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"code": 2, "msg": err.Error()})
		return
	}

	if err := collectionService.Collect(userID, req.PostID); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "msg": "收藏成功"})
}
