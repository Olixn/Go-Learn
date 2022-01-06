/**
 * @Author: Ne-21
 * @Description: 求1-8000素数
 * @File:  case2
 * @Version: 1.0.0
 * @Date: 2022/1/6 15:56
 */

package main

import (
	"fmt"
)

func putNum(initChan chan int) {
	for i := 1; i < 8000; i++ {
		initChan <- i
	}
	fmt.Println("放完了")
	// 放完关闭initChan
	close(initChan)
}

func primeNum(initChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool
	for {

		num, ok := <-initChan
		if !ok {
			// initChan取不到了。
			break
		}
		// 判断素数
		flag = true // 假定传入为素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			// 放入primeChan
			primeChan <- num
		}
	}
	fmt.Println("一个协程取不到数，退出了")
	// 向退出管道exitChan写true
	exitChan <- true
}

func main() {
	initChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	// 开启协程-放数 放入primeChan
	go putNum(initChan)
	// 开启四个协程，判断素数
	for i := 0; i < 4; i++ {
		go primeNum(initChan, primeChan, exitChan)
	}

	// 判断何时关闭primeChan,何时结束主进程
	go func() {
		for i := 0; i < 4; i++ {
			// 取不到四个会阻塞
			<-exitChan
		}
		// 渠道四个则证明写成全部执行完，关闭primeChan
		close(primeChan)
	}()

	// 遍历primeChan，取值，等取不到了，在关闭exitChan
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数为", v)
	}

	fmt.Println("结束")

}
