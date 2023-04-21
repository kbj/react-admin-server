// Package system 角色Vo
// @author: kbj
// @date: 2023/4/12
package system

type RoleSearch struct {
	RoleName *string `json:"roleName,omitempty"`
	RoleKey  *string `json:"roleKey,omitempty"`
	Enabled  *string `json:"enabled,omitempty"`
}

type RoleForm struct {
	ID            *uint   `json:"id"`
	Menus         *[]uint `json:"menus"`
	RoleName      *string `json:"roleName" validate:"required,min=1,max=20" comment:"角色名称"`
	RoleKey       *string `json:"roleKey" validate:"required,min=1,max=100" comment:"权限字符"`
	Enabled       *string `json:"enabled"`
	CheckStrictly *bool   `json:"checkStrictly"`
	OrderNum      *int    `json:"orderNum"`
}
