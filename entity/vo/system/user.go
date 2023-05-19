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
	// 昵称
	NickName string `json:"nickName,omitempty"`
	// 是否启用
	Enabled string `json:"enabled"`
	// 部门ID
	DeptId uint `json:"deptId,omitempty"`
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
	Password string `json:"password" validate:"omitempty,min=8,max=32" comment:"密码"`
	// 部门ID
	DeptId uint `json:"deptId" validate:"numeric" comment:"部门"`
	// 昵称
	NickName string `json:"nickName" validate:"max=100" comment:"用户昵称"`
	// 是否启用
	Enabled string `json:"enabled"`
	// 电子邮箱
	Email string `json:"email" validate:"max=200" comment:"邮箱"`
	// 角色ID
	Roles *[]uint `json:"roles"`
}

type ResetPasswordRequest struct {
	CurrentPassword string `json:"currentPassword" validate:"omitempty,min=8,max=32" comment:"当前密码"`
	NewPassword     string `json:"newPassword" validate:"omitempty,min=8,max=32" comment:"新密码"`
}
