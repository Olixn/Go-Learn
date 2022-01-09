/**
 * @Author: Ne-21
 * @Description: Redis 连接池
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/1/9 10:06
 */

package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 定义一个全局pool
var pool *redis.Pool

// 初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大连接空闲数
		MaxActive:   0,   // 表示和数据库的最大连接数，0为没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	// 从pool取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("Set err = ", err)
		return
	}

	n, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("Get err = ", err)
		return
	}
	fmt.Println("name = ", n)
}
