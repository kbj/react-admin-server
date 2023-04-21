// Package domain 角色
// @author: kbj
// @date: 2023/2/10
package domain

type Role struct {
	Common
	RoleName      *string `json:"roleName" gorm:"size:20;comment:角色名称;not null"`
	RoleKey       *string `json:"roleKey" gorm:"size:100;comment:角色权限字符串;not null;unique"`
	CheckStrictly *bool   `json:"checkStrictly" gorm:"comment:是否父子联动;default:true"`
	Enabled       *string `json:"enabled" gorm:"size:1;comment:是否启用;not null;default:1;type:char"`
	OrderNum      *int    `json:"orderNum" gorm:"comment:排序;not null;default:0"`
	Users         *[]User `json:"users,omitempty" gorm:"many2many:user_role"`
	Menus         *[]Menu `json:"menus,omitempty" gorm:"many2many:role_menu"`
}
