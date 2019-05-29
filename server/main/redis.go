package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func initRedisPool(maxIdle, maxActive int, idleTimeout time.Duration, host string) {
	pool = &redis.Pool{
		// 初始化链接数量
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host)
		},
	}
}
