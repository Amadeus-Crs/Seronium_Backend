package handler

import (
	"Seronium/internal/config"
	"Seronium/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var userService = service.NewUserService()

func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2, "msg": err.Error()})
		return
	}

	if err := userService.Register(req.Username, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": "注册失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "注册成功"})
}

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2, "msg": err.Error()})
		return
	}

	user, err := userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 2, "msg": "登录失败"})
		return
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": "token 生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "data": gin.H{"token": tokenStr}})
}

func GetProfile(c *gin.Context) {
	userID := c.GetUint64("user_id")
	user, err := userService.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": "获取用户信息失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": user})
}

func UpdateProfile(c *gin.Context) {
	userID := c.GetUint64("user_id")
	var req struct {
		Bio       string `json:"bio"`
		AvatarURL string `json:"avatar_url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2, "msg": err.Error()})
		return
	}

	if err := userService.UpdateProfile(userID, req.Bio, req.AvatarURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 2, "msg": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "更新成功"})
}
