/**
 * @Author: Ne-21
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/1/5 17:47
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() bool {
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("序列化失败", err)
		return false
	}

	err = ioutil.WriteFile("./test.txt", data, 0666)
	if err != nil {
		fmt.Println("写入失败", err)
		return false
	}
	return true
}

func (m *Monster) ReStore() bool {
	str, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("读入失败")
		return false
	}

	err = json.Unmarshal(str, &m)
	if err != nil {
		fmt.Println("反序列化失败")
		return false
	}
	return true
}
