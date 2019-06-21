package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiangrui2019/redis"
)

func LimiterMiddleware(serviceName string, client redis.Client, limit int64, duration int32) func(context *gin.Context) {
	limiterservice, err := NewRedisLimiter(client)

	if err != nil {
		panic(err)
	}

	return func(context *gin.Context) {
		err := limiterservice.Limit(serviceName, context.ClientIP()+context.Request.Method+context.Request.RequestURI, limit, duration)

		if err != nil {
			context.AbortWithStatus(400)
			return
		} else {
			context.Next()
		}
	}
}