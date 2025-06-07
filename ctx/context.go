package ctx

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Context struct {
	redis *redis.Client
}

func New() *Context {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("Connected to Redis")

	return &Context{
		redis: rdb,
	}
}

func (ctx *Context) Cancel() {
	ctx.redis.Close()
}

func (ctx *Context) Redis() *redis.Client {
	return ctx.redis
}
