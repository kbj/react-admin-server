// Package domain 菜单
// @author: kbj
// @date: 2023/2/10
package domain

import "react-admin-server/global/types"

type Menu struct {
	Common
	MenuName       string         `json:"menuName" gorm:"size:100;comment:菜单名称;not null"`
	ParentId       uint           `json:"parentId" gorm:"comment:父菜单ID;not null;index;default:0"`
	OrderNum       int            `json:"orderNum" gorm:"comment:排序;not null;default:0"`
	MenuType       types.MenuType `json:"menuType" gorm:"size:1;comment:菜单类型(C目录M菜单B按钮);not null;type:char"`
	Visible        bool           `json:"visible" gorm:"comment:是否可见;not null;default:true"`
	Enabled        bool           `json:"enabled" gorm:"comment:是否启用;not null;default:true"`
	PermissionFlag string         `json:"permissionFlag" gorm:"size:100;comment:权限标识;index"`
	Icon           string         `json:"icon" gorm:"size:100;comment:图标"`
}
