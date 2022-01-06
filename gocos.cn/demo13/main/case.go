/**
 * @Author: Ne-21
 * @Description: goroutine 和 channel 结合案例
 * @File:  case
 * @Version: 1.0.0
 * @Date: 2022/1/6 14:59
 */

package main

import (
	"fmt"
)

func writeData(initChan chan int) {
	for i := 0; i < 50; i++ {
		// 放入数据
		initChan <- i
		fmt.Println("写", i)

	}
	close(initChan)
}

func readData(initChan chan int, exitChan chan bool) {
	for {
		v, ok := <-initChan
		if !ok {
			break
		}
		fmt.Println("读", v)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	// 创建两个管道
	initChan := make(chan int, 50)
	exitChan := make(chan bool, 1) // 用来控制主线程的退出

	// 启动协程
	go writeData(initChan)
	go readData(initChan, exitChan)

	//
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
