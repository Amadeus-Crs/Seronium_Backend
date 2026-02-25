package middleware

import (
	"Seronium/internal/config"
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

var JWTMiddleware *jwt.HertzJWTMiddleware

func InitJWTMiddleware() {
	var err error
	JWTMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "quora",
		Key:           []byte(config.JWTSecret),
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour * 24,
		IdentityKey:   "user_id",
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{
					"user_id": v["user_id"],
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims
		},
	})
	if err != nil {
		panic(err)
	}
}

func GetUserID(ctx context.Context, c *app.RequestContext) uint64 {
	claims := jwt.ExtractClaims(ctx, c)
	if userID, ok := claims["user_id"].(float64); ok {
		return uint64(userID)
	}
	return 0
}
