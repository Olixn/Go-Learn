/**
 * @Author: Ne-21
 * @Description:
 * @File:  exercise
 * @Version: 1.0.0
 * @Date: 2022/1/6 14:37
 */

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var personChan chan Person
	personChan = make(chan Person, 10)

	for i := 0; i < 10; i++ {
		person := Person{
			Name:    "姓名" + strconv.Itoa(rand.Intn(100)),
			Age:     rand.Intn(20),
			Address: "地址" + strconv.Itoa(rand.Intn(100)),
		}

		// 推入
		personChan <- person
	}

	for i := 0; i < 10; i++ {
		p := <-personChan
		fmt.Printf("person%v name = %v, age = %v, address = %v \n", i+1, p.Name, p.Age, p.Address)
	}
}
