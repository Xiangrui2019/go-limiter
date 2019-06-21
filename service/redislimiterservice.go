package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/xiangrui2019/redis"
	"strconv"
)
type RedisLimiterService struct {
	ctx context.Context
	client redis.Client
}

func NewRedisLimiter(client redis.Client) (*RedisLimiterService, error) {
	service := new(RedisLimiterService)
	service.ctx = context.Background()
	service.client = client
	err := service.client.Ping(service.ctx)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (service *RedisLimiterService) Limit(serviceName string, clientid string, limit int64, duration int32) error {
	var sum int64
	var err error

	data, err := service.client.Get(service.ctx, serviceName + "&-&" +clientid)

	if data == nil {
		sum = 1
		sumstring := strconv.Itoa(int(sum))

		err := service.client.Set(service.ctx, &redis.Item{
			Key:   serviceName + "&-&" + clientid,
			Value: []byte(sumstring),
			TTL:   duration,
		})

		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	} else {
		sum, err = strconv.ParseInt(string(data.Value), 10, 0)

		if err != nil {
			return err
		}

		if sum >= limit {
			return errors.New("limit too big.")
		} else {
			sum, err = service.client.IncrBy(service.ctx, serviceName + "&-&" + clientid, 1)

			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}