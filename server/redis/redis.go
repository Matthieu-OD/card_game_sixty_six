package redisDB

import (
	"Matthieu-OD/card_game_sixty_six/server/game"

	"github.com/redis/go-redis/v9"

	"context"
	"fmt"
	"log"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := checkConnection(c)

	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		fmt.Println("Connected to Redis")
	}

	return c
}

func checkConnection(c *redis.Client) error {
	pong, err := c.Ping(ctx).Result()
	fmt.Println(pong, err)
	return err
}

func StoreGameid(c *redis.Client, id string) error {
	err := c.Set(ctx, "game-id", id, 0).Err()
	return err
}

func StoreGame(c *redis.Client, game game.Game) error {
	return nil
}

func DeleteGameid(c *redis.Client, id string) error {
	_, err := c.Del(ctx, id).Result()
	return err
}

func IdExists(c *redis.Client, id string) (bool, error) {
	exists, err := c.Exists(ctx, id).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

// TODO: add function to add a new party data
// TODO: add function to get the party data
