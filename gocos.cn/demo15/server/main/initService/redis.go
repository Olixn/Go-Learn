/**
 * @Author: Ne-21
 * @Description:
 * @File:  redis
 * @Version: 1.0.0
 * @Date: 2022/1/10 15:03
 */

package initService

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func InitRedisPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大连接空闲数
		MaxActive:   maxActive,   // 表示和数据库的最大连接数，0为没有限制
		IdleTimeout: idleTimeout, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接
			return redis.Dial("tcp", address)
		},
	}
}
