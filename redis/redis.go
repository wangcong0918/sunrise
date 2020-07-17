package redis

import "github.com/go-redis/redis"

type Config struct {
	Options redis.Options
}

func (c Config) NewRedisClient() (client *redis.Client, err error) {
	client = redis.NewClient(&c.Options)

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return client, nil
}
