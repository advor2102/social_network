package repository

import "github.com/redis/go-redis/v9"

type Cache struct {
	rdb *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{
		rdb: client,
	}
}
