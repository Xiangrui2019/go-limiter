package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiangrui2019/go-limiter/service"
	"github.com/xiangrui2019/redis"
)

func LimiterMiddleware(client redis.Client) func(context *gin.Context) {
	limiterservice := service.NewRedisLimiter(client)

	return func(context *gin.Context) {

	}
}