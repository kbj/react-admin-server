// Package service 全局Service出口
// @author: kbj
// @date: 2023/2/3
package service

import "react-admin-server/service/system"

var (
	LoginService = new(system.LoginService)
	UserService  = new(system.UserService)
	DictService  = new(system.DictService)
)
