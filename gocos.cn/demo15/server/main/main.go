/**
 * @Author: Ne-21
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/1/9 16:42
 */

package main

import (
	"fmt"
	"gocos.cn/demo15/server/controller"
	"net"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 延时关闭
	defer conn.Close()

	// 调用主控制器
	processor := &controller.Processor{
		Conn: conn,
	}
	err := processor.Process()
	if err != nil {
		fmt.Println("客户端与服务器端通信协程错误 = ", err)
		return
	}

}

func main() {
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listen.Close()

	// 等待客户端链接
	for {
		fmt.Println("等待客户端链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.accept() err = ", err)
		}

		// 起协程处理，和客户端保持通讯
		go process(conn)
	}
}
