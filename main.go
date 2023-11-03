package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use the appropriate address
		Password: "",               // no password set by default
		DB:       0,                // use default DB
	})

	// Check the connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong) // Output: PONG

	// Set a key
	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// Get a key
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key:", val)

	// Update a key (using set to update the value of the existing key)
	err = rdb.Set(ctx, "key", "new_value", 0).Err()
	if err != nil {
		panic(err)
	}

	// Delete a key
	err = rdb.Del(ctx, "key").Err()
	if err != nil {
		panic(err)
	}

	// Optionally, close the connection when your application exits
	err = rdb.Close()
	if err != nil {
		panic(err)
	}
}
