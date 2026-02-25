package main

import (
	"Seronium/internal/config"
	"Seronium/internal/handler"
	"Seronium/internal/middleware"
	"Seronium/internal/repository"
	"Seronium/internal/util"
	"log"

	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
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

	h := server.Default(server.WithHostPorts(":8080"))
	h.Use(recovery.Recovery())
	h.Use(middleware.CORS())
	middleware.InitJWTMiddleware()

	authGroup := h.Group("/api/auth")
	{
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/login", handler.Login)
	}

	apiGroup := h.Group("/api")
	apiGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
	{
		apiGroup.GET("/user/profile", handler.GetProfile)
		apiGroup.PUT("/user/profile", handler.UpdateProfile)

		apiGroup.POST("/posts", handler.CreatePost)
		apiGroup.GET("/posts/:id", handler.GetPost)
		apiGroup.PUT("/posts/:id", handler.UpdatePost)
		apiGroup.DELETE("/posts/:id", handler.DeletePost)
		apiGroup.GET("/posts", handler.ListPosts)

		apiGroup.POST("/comments", handler.CreateComment)
		apiGroup.POST("/likes", handler.Like)
		apiGroup.POST("/collections", handler.Collect)
		apiGroup.POST("/upload", handler.Image)
	}

	h.Spin()
}
