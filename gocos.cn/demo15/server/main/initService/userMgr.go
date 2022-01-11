/**
 * @Author: Ne-21
 * @Description:
 * @File:  userMgr
 * @Version: 1.0.0
 * @Date: 2022/1/11 14:15
 */

package initService

import (
	"gocos.cn/demo15/server/controller"
)

func InitUserMgr() {
	controller.MyUserMgr = &controller.UserMgr{
		OnlineUsers: make(map[int]*controller.UserProcess, 1024),
	}
}
