/**
 * @Author: Ne-21
 * @Description: 反射
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/1/7 13:48
 */

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTest(i interface{}) {
	// Type 是类型，Kind 是类别，二者可能相同，也可能不同
	// TypeOf返回接口中保存的值的类型
	rTyp := reflect.TypeOf(i)
	fmt.Println("rType = ", rTyp) // rType =  main.Student

	// ValueOf返回一个初始化为i接口保管的具体值的Value
	rVal := reflect.ValueOf(i)
	fmt.Println("rVal = ", rVal) // rVal =  {小白 20}

	// 获取 变量对应的 kind
	fmt.Println("kind = ", rVal.Kind()) // kind =  struct
	fmt.Println("kind = ", rTyp.Kind()) // kind =  struct

	// 将 reflect.value 转化为 Interface{}
	iv := rVal.Interface()
	// 类型断言
	fmt.Println("stu.Name = ", iv.(Student).Name)

}

func reflectTest02(i interface{}) {
	// 1.通过反射修改 num int 的值
	rVal := reflect.ValueOf(i)
	// 当传入地址时，rVal不再是reflect.value 而要通过 Elem() 方法
	// Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。如果v的Kind不是Interface或Ptr会panic
	fmt.Println("rVal kind = ", rVal.Kind())               // rVal kind =  ptr
	fmt.Println("rVal.Elem() kind = ", rVal.Elem().Kind()) // rVal.Elem() kind =  int
	// 修改num int  rVal.Elem() 可以理解为 *num
	rVal.Elem().SetInt(20)
}

func main() {
	stu := Student{"小白", 20}
	reflectTest(stu)

	var num int = 10
	// 传入地址
	reflectTest02(&num)
	fmt.Println(num)
}
