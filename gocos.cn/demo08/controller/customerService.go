/**
 * @Author: Ne-21
 * @Description:
 * @File:  customerServes
 * @Version: 1.0.0
 * @Date: 2022/1/4 9:45
 */

package controller

import "gocos.cn/demo08/model"

// CustomerService 完成对Customer的操作 CRUD
type CustomerService struct {
	customers []model.Customer
	// 声明一个字段，表示当前切片含有多少个用户
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "1234567890", "123456@163.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

func (cs *CustomerService) Add(customer model.Customer) bool {
	// 分配id
	customer.Id = cs.customerNum + 1
	cs.customers = append(cs.customers, customer)
	return true
}

func (cs *CustomerService) FindById(id int) int {
	for i, customer := range cs.customers {
		if customer.Id == id {
			return i
		}
	}
	return -1
}

func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	} else {
		cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
		return true
	}
}

func (cs *CustomerService) Update(customer model.Customer) bool {
	index := cs.FindById(customer.Id)
	if index == -1 {
		return false
	}
	if customer.Name != "" {
		cs.customers[index].Name = customer.Name
	}
	if customer.Age != 0 {
		cs.customers[index].Age = customer.Age
	}
	if customer.Gender != "" {
		cs.customers[index].Gender = customer.Gender
	}
	if customer.Phone != "" {
		cs.customers[index].Phone = customer.Phone
	}
	if customer.Email != "" {
		cs.customers[index].Email = customer.Email
	}
	return true

}
