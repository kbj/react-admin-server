// Package domain 用户
// @author: kbj
// @date: 2023/2/10
package domain

import "react-admin-server/global/types"

type User struct {
	Common
	Username string       `json:"username,omitempty" gorm:"size:32;comment:用户名;<-:create;unique;not null"`
	Password string       `json:"-" gorm:"comment:密码;not null"`
	DeptId   uint         `json:"deptId,omitempty" gorm:"comment:部门ID"`
	Mobile   string       `json:"mobile,omitempty" gorm:"comment:手机号;index;size:20"`
	Gender   types.Gender `json:"gender,omitempty" gorm:"comment:性别;type:char;size:1"`
	Avatar   string       `json:"avatar,omitempty" gorm:"comment:头像;size:200"`
	Roles    []*Role      `json:"-" gorm:"many2many:user_role"`
}
