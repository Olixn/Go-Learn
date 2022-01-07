/**
 * @Author: Ne-21
 * @Description: 练习
 * @File:  example2
 * @Version: 1.0.0
 * @Date: 2022/1/7 15:28
 */

package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%v 完成了减法运行 %v - %v = %v \n", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func reflectTest03(a interface{}) {
	rVal := reflect.ValueOf(a).Elem()

	num := rVal.NumField()
	for i := 0; i < num; i++ {
		fmt.Printf("字段%v是%v \n", i, rVal.Field(i))
	}

	rVal.Field(0).SetInt(8)
	rVal.Field(1).SetInt(3)

	rVal.Method(0).Call([]reflect.Value{reflect.ValueOf("小白")})

}

func main() {
	var a = Cal{}
	reflectTest03(&a)
}
