/**
 * @Author: Ne-21
 * @Description: 全局变量的互斥锁
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/1/6 12:04
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int)
	lock  sync.Mutex // 全局互斥锁
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}

func main() {
	for i := 1; i < 20; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 5)
	lock.Lock()
	for i, v := range myMap {
		fmt.Println(i, v)
	}
	lock.Unlock()
}
