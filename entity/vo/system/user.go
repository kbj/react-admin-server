// Package system
// @author: kbj
// @date: 2023/2/9
package system

import (
	"react-admin-server/global/types"
)

type UserSearch struct {
	// 用户名
	Username string `json:"username,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty"`
	// 性别
	Gender types.Gender `json:"gender,omitempty"`
}

type UserRequest struct {
	ID int `json:"id"`
	// 用户名
	Username string `json:"username" validate:"required,min=5,max=16" comment:"用户名"`
	// 手机号
	Mobile string `json:"mobile" validate:"max=20" comment:"手机号"`
	// 性别
	Gender types.Gender `json:"gender" validate:"max=1" comment:"性别"`
	// 密码
	Password string `json:"password" validate:"required,min=8,max=64" comment:"密码"`
	// 部门ID
	DeptId uint `json:"deptId" validate:"numeric" comment:"部门"`
	// 角色ID
	Roles *[]uint `json:"roles"`
}
