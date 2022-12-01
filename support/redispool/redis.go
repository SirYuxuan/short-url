package redispool

import "github.com/gomodule/redigo/redis"

var Pool *redis.Pool

func Init() {
	Pool = &redis.Pool{MaxIdle: 8, MaxActive: 0, IdleTimeout: 100, Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(1))
	}}
}
