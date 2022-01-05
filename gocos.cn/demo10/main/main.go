/**
 * @Author: Ne-21
 * @Description: 反序列化
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/1/5 16:50
 */

package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func unmarshalStruct() {
	str := `{"name":"mike","age":0,"birthday":"2011-11-11","sal":8000,"skill":"吃饭"}`

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化失败")
	}

	fmt.Println(monster)
}

func unmarshalMap() {
	str := `{"name":"mike","age":0,"birthday":"2011-11-11","sal":8000,"skill":"吃饭"}`

	var a map[string]interface{}
	// 反序列化不需要make 已经封装了
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("反序列化失败")
	}

	fmt.Println(a)
}

func unmarshalSlice() {
	str := `[{"address":"东北","age":null,"name":"sli"},{"address":["东北","北京"],"age":88,"name":"ddg"}]`

	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("反序列化失败")
	}

	fmt.Println(slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
