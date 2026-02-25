package handler

import (
	"Seronium/internal/config"
	"Seronium/internal/middleware"
	"Seronium/internal/service"
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt"
)

var userService = service.NewUserService()

func Register(ctx context.Context, c *app.RequestContext) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"code": 2, "msg": err.Error()})
		return
	}

	if err := userService.Register(req.Username, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": "注册失败"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "msg": "注册成功"})
}

func Login(ctx context.Context, c *app.RequestContext) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"code": 2, "msg": err.Error()})
		return
	}

	user, err := userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"code": 2, "msg": "用户名或密码错误"})
		return
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": "token 生成失败"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 1,
		"msg":  "登录成功",
		"data": map[string]interface{}{
			"token": tokenStr,
		},
	})
}

func GetProfile(ctx context.Context, c *app.RequestContext) {
	userID := middleware.GetUserID(ctx, c)
	user, err := userService.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": "获取用户信息失败"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "data": user})
}

func UpdateProfile(ctx context.Context, c *app.RequestContext) {
	userID := middleware.GetUserID(ctx, c)
	var req struct {
		Bio       string `json:"bio"`
		AvatarURL string `json:"avatar_url"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"code": 2, "msg": err.Error()})
		return
	}

	if err := userService.UpdateProfile(userID, req.Bio, req.AvatarURL); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 2, "msg": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "msg": "更新成功"})
}
