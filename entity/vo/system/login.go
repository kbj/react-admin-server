// Package system	登录部分的接口VO
// @author: kbj
// @date: 2023/2/2
package system

import (
	"react-admin-server/global/types"
)

type LoginRequest struct {
	Username string `json:"username,omitempty" validate:"required,min=5,max=16" comment:"用户名"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=64" comment:"密码"`
}

type LoginUserResponse struct {
	ID       uint         `json:"id,omitempty"`
	CreateAt int64        `json:"createAt,omitempty"`
	UpdateAt int64        `json:"updateAt,omitempty"`
	Username string       `json:"username,omitempty"`
	Mobile   string       `json:"mobile,omitempty"`
	Gender   types.Gender `json:"gender,omitempty"`
	Avatar   string       `json:"avatar,omitempty"`
	DeptId   uint         `json:"deptId,omitempty"`
	NickName string       `json:"nickName,omitempty"`
	Email    string       `json:"email,omitempty"`
}
