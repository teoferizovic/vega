package database

import (
	"github.com/go-redis/redis/v8"
)

// DBCon is the connection handle for Redis
var (
	Redis *redis.Client
)
