package handler

import (
	"Seronium/internal/service"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

var uploadService = service.NewUploadService()

func Image(ctx context.Context, c *app.RequestContext) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"code": 2, "msg": "请选择文件"})
		return
	}

	url, err := uploadService.UploadImage(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": "上传失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 1,
		"msg":  "上传成功",
		"data": map[string]interface{}{"url": url},
	})
}
