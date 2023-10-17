package redis

import (
	"fmt"
	"log"
	"strconv"
	"context"
	"github.com/dimasyotama/go-notes/config"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	var err error
	p := config.Config("REDIS_PORT")
	redis_port, err := strconv.ParseUint(p, 10, 32)
	redis_addr := config.Config("REDIS_ADDRESS")
	redis_pwd := config.Config("REDIS_PASSWORD")

	// Create the Redis server address by combining address and port
	redisServerAddress := fmt.Sprintf("%s:%d", redis_addr, redis_port)
	
	redis_connect := redis.NewClient(&redis.Options{
		Addr: 		redisServerAddress,
		Password: 	redis_pwd,
		DB: 		0,
	})

	// Ping the Redis server to check the connection
	if err := redis_connect.Ping(context.Background()).Err(); err != nil {
		fmt.Printf("Failed to connect to Redis: %v", err)
	}

	if err != nil{
		log.Panic("Failed to connect redis")
	}
	fmt.Println("Connection to Redis Successfully")
	
	return redis_connect


}