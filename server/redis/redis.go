package redidsclient

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return c
}

// TODO: add function to add a new party data
// TODO: add function to get the party data
