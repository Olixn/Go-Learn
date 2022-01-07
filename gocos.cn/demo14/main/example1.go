/**
 * @Author: Ne-21
 * @Description: 反射的最佳实践
 * @File:  example1
 * @Version: 1.0.0
 * @Date: 2022/1/7 14:38
 */

package main

import (
	"fmt"
	"reflect"
)

// 使用反射来遍历结构体字段，调用结构体的方法，并获取结构体标签的值

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (m Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(m)
	fmt.Println("---end---")
}

func (m Monster) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

func (m Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func TestStruct(a interface{}) {
	// 获取reflect.Type 类型
	rTyp := reflect.TypeOf(a).Elem() // 传入的为地址，要Elem()
	// 获取reflect.Value 类型
	rVal := reflect.ValueOf(a).Elem()
	// 获取到a对应的类别
	kd := rVal.Kind()
	// 判断传入的是不是结构体
	if kd != reflect.Struct {
		fmt.Println("不是struct")
		return
	}

	// 获取到该结构的字段数量
	num := rVal.NumField()
	fmt.Println("该结构体的字段数量是", num)
	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("字段 = %v 值 = %v \n", i, rVal.Field(i))
		// 获取结构体的标签，注意是通过 reflect.Type 来获取tag标签的值
		tagVal := rTyp.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("字段 = %v tag值 = %v \n", i, tagVal)
		}
	}

	// 获取该结构体有多少个方法
	numOfMethod := rVal.NumMethod()
	fmt.Println("该结构体的方法数量是", numOfMethod)

	// 方法的排序是按照函数名排序（ASCII码）
	rVal.Method(1).Call(nil) // 调用了Print()

	// 调用GetSum()
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := rVal.Method(0).Call(params)  // 传入[]reflect.Value
	fmt.Println("res = ", res[0].Int()) // 返回[]reflect.Value

	// 修改字段
	rVal.Field(0).SetString("小红")

}

func main() {
	var a Monster = Monster{
		Name:  "小白",
		Age:   40,
		Score: 30.8,
	}
	TestStruct(&a)
	fmt.Println(a)
}
