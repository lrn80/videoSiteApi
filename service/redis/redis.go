package redisClient

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
)

func Connect() redis.Conn {
	pool, _ := redis.Dial("tcp", "47.106.252.126:6379")
	return pool
}

func PoolConnect() redis.Conn  {
   pool :=	&redis.Pool{
		MaxIdle: 1,
		MaxActive: 10,
		IdleTimeout: 180 * time.Second,
		Wait: true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.String(
				"redisdb" + ":" + beego.AppConfig.String("redisport")))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	return pool.Get()
}