package ratelimiter

import "linkpulse/internal/repository/redis"


type Repository struct{
	redis redis.Redis
}

func New (redisRepo redis.Redis) Repository {
	return Repository{
		redis: redisRepo,
	}
}