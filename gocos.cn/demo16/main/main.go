/**
 * @Author: Ne-21
 * @Description: 操作Redis
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/1/8 17:40
 */

package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dail err = ", err)
		return
	}
	defer conn.Close()

	// 写入数据 string [key-value]
	_, err = conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("Set err = ", err)
		return
	}

	// 读取数据 string [key-value]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err = ", err)
	}

	fmt.Println(r)

	// 写入数据 hash
	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("HSet err = ", err)
		return
	}
	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("HSet err = ", err)
		return
	}

	// 读取数据 hash
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("HGet err = ", err)
	}

	r2, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("HGet err = ", err)
	}

	fmt.Println(r1, r2)

	// 写入数据 List
	_, err = conn.Do("LPush", "heroList", "no1:宋江", 18, "no2:卢俊义", 20)

	// 读取数据 List
	r, err = redis.String(conn.Do("RPop", "heroList"))

	fmt.Println(r)

	// 批量操作
	_, err = redis.String(conn.Do("HMSet", "user02", "name", "小白", "age", 30))
	if err != nil {
		fmt.Println("HMSet err = ", err)
	}

	r3, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("HMGet err = ", err)
	}
	for _, v := range r3 {
		fmt.Println(v)
	}

	// 给 name 设置有效时间
	_, err = conn.Do("set", "name01", "mike")
	_, err = conn.Do("expire", "name01", 10)
	r4, _ := redis.String(conn.Do("Get", "name01"))
	fmt.Println(r4)
	time.Sleep(time.Second * 11)
	r5, _ := redis.String(conn.Do("Get", "name01"))
	fmt.Println("11秒后的name01 = ", r5)
}
