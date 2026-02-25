package handler

import (
	"Seronium/internal/middleware"
	"Seronium/internal/service"
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

var postService = service.NewPostService()

func CreatePost(ctx context.Context, c *app.RequestContext) {
	userID := middleware.GetUserID(ctx, c)
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Type    string `json:"type" binding:"required"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"code": 2, "msg": err.Error()})
		return
	}

	id, err := postService.Create(userID, req.Title, req.Content, req.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": "创建失败"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "data": map[string]interface{}{"post_id": id}})
}

func GetPost(ctx context.Context, c *app.RequestContext) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	post, err := postService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"code": 2, "msg": "未找到"})
		return
	}
	postService.IncView(id)
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "data": post})
}

func UpdatePost(ctx context.Context, c *app.RequestContext) {
	userID := middleware.GetUserID(ctx, c)
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"code": 2, "msg": err.Error()})
		return
	}

	if err := postService.Update(userID, id, req.Title, req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "msg": "更新成功"})
}

func DeletePost(ctx context.Context, c *app.RequestContext) {
	userID := middleware.GetUserID(ctx, c)
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	if err := postService.Delete(userID, id); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "msg": "删除成功"})
}

func ListPosts(ctx context.Context, c *app.RequestContext) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	sort := c.Query("sort")

	posts, err := postService.List(offset, limit, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": "列表失败"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "data": posts})
}
