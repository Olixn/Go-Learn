/**
 * @Author: Ne-21
 * @Description:
 * @File:  customerView
 * @Version: 1.0.0
 * @Date: 2022/1/4 9:38
 */

package view

import (
	"fmt"
	"gocos.cn/demo08/controller"
	"gocos.cn/demo08/model"
)

type CustomerView struct {
	// 定义必要的字段
	key             string
	loop            bool
	customerService *controller.CustomerService
}

// NewCustomerView 返回并初始化一个CustomerView实例
func NewCustomerView() CustomerView {
	return CustomerView{
		key:             "",
		loop:            true,
		customerService: controller.NewCustomerService(),
	}
}

// 显示所有的客户信息
func (cv *CustomerView) list() {
	customers := cv.customerService.List()
	fmt.Println("----------------客户列表---------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, customer := range customers {
		fmt.Println(customer.GetInfo())
	}
	fmt.Println("--------------客户列表结束--------------")
}

// 得到用户的输入信息，构建新的客户，并完成添加
func (cv *CustomerView) add() {
	fmt.Println("----------------添加客户---------------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("E-Mail")
	email := ""
	fmt.Scanln(&email)

	// 构建新的Customer实例
	customer := model.NewCustomerNoId(name, gender, age, phone, email)
	//调用添加
	if cv.customerService.Add(customer) {
		fmt.Println("----------------添加成功---------------")
	} else {
		fmt.Println("----------------添加失败---------------")
	}

}

// 删除id对应的客户
func (cv *CustomerView) delete() {
	fmt.Println("----------------删除客户---------------")
	fmt.Println("请选择待删除的客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除（Y/N）：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if cv.customerService.Delete(id) {
			fmt.Println("----------------删除完成---------------")
		} else {
			fmt.Println("---------删除失败，用户编号不存在---------")
		}
	}

}

func (cv *CustomerView) exit() {
	fmt.Println("确定退出吗（Y/N）")
	ifExit := ""
	fmt.Scanln(&ifExit)
	if ifExit == "y" || ifExit == "Y" {
		cv.loop = false
		fmt.Println("退出成功")
	}
}

func (cv *CustomerView) update() {
	fmt.Println("----------------修改客户---------------")
	fmt.Println("请输入待修改的用户编号（-1退出）：")
	key := -1
	fmt.Scanln(&key)
	if key == -1 {
		return
	} else {
		index := cv.customerService.FindById(key)
		if index == -1 {
			fmt.Println("用户不存在，请重新输入")
			cv.update()
		}
		customers := cv.customerService.List()
		fmt.Printf("姓名(%v) \n", customers[index].Name)
		name := ""
		fmt.Scanln(&name)
		fmt.Printf("性别(%v) \n", customers[index].Gender)
		gender := ""
		fmt.Scanln(&gender)
		fmt.Printf("年龄(%v) \n", customers[index].Age)
		age := 0
		fmt.Scanln(&age)
		fmt.Printf("电话(%v) \n", customers[index].Phone)
		phone := ""
		fmt.Scanln(&phone)
		fmt.Printf("E-mail(%v) \n", customers[index].Email)
		email := ""
		fmt.Scanln(&email)
		if cv.customerService.Update(model.NewCustomer(customers[index].Id, name, gender, age, phone, email)) {
			fmt.Println("----------------修改成功---------------")
		}
	}
}

// MainMenu 显示主菜单
func (cv *CustomerView) MainMenu() {
	for {
		fmt.Println("--------------客户信息管理系统-------------")
		fmt.Println("              1 添加客户")
		fmt.Println("              2 修改客户")
		fmt.Println("              3 删除客户")
		fmt.Println("              4 客户列表")
		fmt.Println("              5 退   出")
		fmt.Println("请选择（1-5）：")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("输入有误，请重新输入")
		}

		if !cv.loop {
			break
		}
	}

}
