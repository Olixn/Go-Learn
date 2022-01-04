/**
 * @Author: Ne-21
 * @Description:
 * @File:  customer
 * @Version: 1.0.0
 * @Date: 2022/1/4 9:36
 */

package model

import "fmt"

// Customer 声明一个Customer结构体,表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// NewCustomer 使用工厂模式，返回一个Customer实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomerNoId(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// GetInfo 返回当前用户信息
func (c Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return info
}
