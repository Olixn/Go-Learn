/**
 * @Author: Ne-21
 * @Description: 序列化
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/1/5 16:01
 */

package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name,omitempty"` // omitempty 关键字，来表示这条信息如果没有提供，在序列化成 json 的时候就不要包含其默认值。
	Age      int     `json:"age,omitempty"`
	Birthday string  `json:"birthday,omitempty"`
	Sal      float64 `json:"sal,omitempty"`
	Skill    string  `json:"skill,omitempty"`
}

func testStruct() {
	monster := Monster{
		Name:     "mike",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "吃饭",
	}

	// 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	fmt.Println(string(data))
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "join"
	a["age"] = 30
	a["address"] = "beijing"

	// 序列化
	data, err := json.Marshal(&a)
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	fmt.Println(string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "sli"
	m1["age"] = nil // null
	m1["address"] = "东北"

	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "ddg"
	m2["age"] = 88
	m2["address"] = [2]string{"东北", "北京"}

	slice = append(slice, m2)

	// 序列化
	data, err := json.Marshal(&slice)
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	fmt.Println(string(data))
}

func testFloat64() {
	var num1 float64 = 2345.36

	// 序列化
	data, err := json.Marshal(&num1)
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	fmt.Println(string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
