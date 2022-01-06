/**
 * @Author: Ne-21
 * @Description: channel管道
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/1/6 14:22
 */

package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// 声明一个管道
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	cat1 := Cat{"小花", 10}
	cat2 := Cat{"小白", 66}

	var a map[string]interface{}
	a = make(map[string]interface{})

	a["name"] = "mike"
	a["age"] = 16

	allChan <- cat1
	allChan <- cat2
	allChan <- a
	// close(allChan) // 关闭channel，在此之后不可以进行写入数据到channel,读取数据正常
	allChan <- 10
	allChan <- "son"

	// 推出第一个元素
	<-allChan

	newCat := <-allChan // 取出第二个元素

	fmt.Printf("newCat type = %T  newCat = %v \n", newCat, newCat)
	// newCat type = main.Cat  newCat = {小白 66}

	// fmt.Printf("newCat.Name = %v", newCat.Name) // 类型错误

	// 类型断言来解决
	fmt.Printf("newCat.Name = %v \n", newCat.(Cat).Name)
	fmt.Println("------------------------------")

	// 遍历channel 用for-range
	// 1 如果管道未关闭，遍历完会报错 fatal error: all goroutines are asleep - deadlock!
	close(allChan)
	// 2 如果管道关闭，则遍历完退出
	for v := range allChan {
		fmt.Println(v)
	}

}
