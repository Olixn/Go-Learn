/**
 * @Author: Ne-21
 * @Description: 公共函数
 * @File:  utils
 * @Version: 1.0.0
 * @Date: 2022/1/9 19:15
 */

package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"gocos.cn/demo15/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte // 传输是使用缓冲
}

func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	// 发送数据data 先发长度，再发data
	// data长度->表示长度的切片
	// buf := make([]byte, 8096)
	fmt.Println("读取数据。。。")
	// conn.Read 只有在conn没有被关闭的情况下，才会阻塞，如果客户端关闭conn则不会阻塞了
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}
	// 根据buf[:4] 转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[:4])

	// 根据 pkgLen 读取消息内容
	n, err := t.Conn.Read(t.Buf[:pkgLen]) // 这里会发生阻塞，直到收到发送的data
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	// 把 pkgLen 反序列化
	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	// 先发送长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(t.Buf[0:4], pkgLen)
	n, err := t.Conn.Write(t.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) err = ", err)
		return
	}
	// 发送 data 本身
	n, err = t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}
	return
}
