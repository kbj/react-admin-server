// Package domain 角色
// @author: kbj
// @date: 2023/2/10
package domain

type Role struct {
	Common
	RoleName string  `json:"roleName" gorm:"size:20;comment:角色名称;not null"`
	RoleKey  string  `json:"roleKey" gorm:"size:100;comment:角色权限字符串;not null;unique"`
	Enabled  bool    `json:"enabled" gorm:"comment:是否启用;not null;default:true"`
	Users    []*User `gorm:"many2many:user_role"`
	Menus    []*Menu `gorm:"many2many:role_menu"`
}
