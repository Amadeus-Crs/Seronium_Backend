package handler

import (
	"Seronium/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var postService = service.NewPostService()

func CreatePost(c *gin.Context) {
	userID := c.GetUint64("user_id")
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Type    string `json:"type" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2, "msg": err.Error()})
		return
	}

	id, err := postService.Create(userID, req.Title, req.Content, req.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": "创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": gin.H{"post_id": id}})
}

func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	post, err := postService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 2, "msg": "未找到"})
		return
	}
	postService.IncView(id)
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": post})
}

func UpdatePost(c *gin.Context) {
	userID := c.GetUint64("user_id")
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2, "msg": err.Error()})
		return
	}

	if err := postService.Update(userID, id, req.Title, req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "更新成功"})
}

func DeletePost(c *gin.Context) {
	userID := c.GetUint64("user_id")
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	if err := postService.Delete(userID, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "删除成功"})
}

func ListPosts(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	sort := c.Query("sort")

	posts, err := postService.List(offset, limit, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": "列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": posts})
}
