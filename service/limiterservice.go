package service

import (
	"context"
	"github.com/xiangrui2019/redis"
)
type RedisLimiterService struct {
	serviceName string
	ctx context.Context
	client redis.Client
}

func (service *LimiterService) Limit() error {

}