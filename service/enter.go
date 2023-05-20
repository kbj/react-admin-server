// Package service 全局Service出口
// @author: kbj
// @date: 2023/2/3
package service

import (
	"react-admin-server/service/common"
	"react-admin-server/service/home"
	"react-admin-server/service/system"
)

var (
	LoginService = new(common.LoginService)
)

var (
	UserService = new(system.UserService)
	DictService = new(system.DictService)
	DeptService = new(system.DeptService)
	MenuService = new(system.MenuService)
	RoleService = new(system.RoleService)
)

var (
	WaterService = new(home.WaterService)
)
