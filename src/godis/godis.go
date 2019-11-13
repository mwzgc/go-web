package godis

import (
	"github.com/go-redis/redis/v7"
)

var client *redis.Client = nil

func getRedis() *redis.Client {
	if client != nil {
		return client
	}

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

// SetValue ...
func SetValue(key, value string) {
	client := getRedis()

	err := client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}

}

// GetValue ...
func GetValue(key string) string {
	client := getRedis()
	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
