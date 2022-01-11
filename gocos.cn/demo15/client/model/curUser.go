/**
 * @Author: Ne-21
 * @Description:
 * @File:  curUser
 * @Version: 1.0.0
 * @Date: 2022/1/11 17:04
 */

package model

import (
	"gocos.cn/demo15/common/message"
	"net"
)

// 将其作为全局，方便调用 （userMgr中进行）

type CurUser struct {
	Conn net.Conn
	message.User
}
