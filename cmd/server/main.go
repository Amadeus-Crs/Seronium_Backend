package main

import (
	"Seronium/internal/config"
	"Seronium/internal/handler"
	"Seronium/internal/middleware"
	"Seronium/internal/repository"
	"Seronium/internal/util"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	util.InitLogger()

	if err := repository.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	repository.InitRedis()

	if err := repository.AutoMigrate(); err != nil {
		zap.L().Warn("auto migrate failed", zap.Error(err))
	}

	repository.InitMinIO()

	r := gin.Default()

	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/login", handler.Login)
	}

	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.JWTAuth())
	{
		apiGroup.GET("/user/profile", handler.GetProfile)
		apiGroup.PUT("/user/profile", handler.UpdateProfile)

		apiGroup.POST("/post", handler.CreatePost)
		apiGroup.GET("/post/:id", handler.GetPost)
		apiGroup.PUT("/posts/:id", handler.UpdatePost)
		apiGroup.DELETE("/posts/:id", handler.DeletePost)
		apiGroup.GET("/posts", handler.ListPosts)

		apiGroup.POST("/comments", handler.CreateComment)
		apiGroup.POST("/likes", handler.Like)
		apiGroup.POST("/collections", handler.Collect)
	}

	if err := r.Run(":8080"); err != nil {
		zap.L().Fatal("failed to run server", zap.Error(err))
	}
}
