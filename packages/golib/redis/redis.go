package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)



type RedisClient struct{
	client *redis.Client
}


func NewConnection() *RedisClient{
	client:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisClient{
		client: client,
	}
}


func (r *RedisClient)PublishMessages(ctx context.Context,channel string,data []byte) error{
	publish:=r.client.Publish(ctx,channel,data).Err()
	return publish
}


func(r *RedisClient)SubscribeMessages(ctx context.Context,channel string)*redis.PubSub {
	subscribe:=r.client.Subscribe(ctx,channel)
	return subscribe
}







