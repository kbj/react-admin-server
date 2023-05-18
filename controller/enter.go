// Package controller
// @author: kbj
// @date: 2023/2/2
package controller

import (
	"react-admin-server/controller/common"
	"react-admin-server/controller/system"
)

var (
	LoginController = new(common.LoginController)
	FileController  = new(common.FileController)
	UserController  = new(system.UserController)
	DictController  = new(system.DictController)
	DeptController  = new(system.DeptController)
	MenuController  = new(system.MenuController)
	RoleController  = new(system.RoleController)
)
