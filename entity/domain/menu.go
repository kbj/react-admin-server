// Package domain 菜单
// @author: kbj
// @date: 2023/2/10
package domain

import "react-admin-server/global/types"

type Menu struct {
	Common
	MenuName       string         `json:"menuName" validate:"required,min=1,max=100" comment:"菜单名称" gorm:"size:100;comment:菜单名称;not null"`
	ParentId       uint           `json:"parentId" gorm:"comment:父菜单ID;not null;index;default:0"`
	OrderNum       int            `json:"orderNum" gorm:"comment:排序;not null;default:0"`
	MenuType       types.MenuType `json:"menuType" validate:"required" comment:"菜单类型" gorm:"size:1;comment:菜单类型(C目录M菜单B按钮);not null;type:char"`
	IsFrame        bool           `json:"isFrame" gorm:"comment:是否外链;default:false"`
	Path           string         `json:"path" gorm:"size:200;comment:路由地址"`
	Component      string         `json:"component" gorm:"size:200;comment:组件地址"`
	Visible        bool           `json:"visible" gorm:"comment:是否可见;not null;default:true"`
	Enabled        string         `json:"enabled" gorm:"comment:是否启用;not null;default:1;size:1;type:char"`
	PermissionFlag string         `json:"permissionFlag" gorm:"size:100;comment:权限标识;index"`
	Icon           string         `json:"icon" gorm:"size:100;comment:图标"`
	Query          string         `json:"query" gorm:"size:200;comment:路由参数"`
	IsCache        bool           `json:"isCache" gorm:"comment:是否缓存;default:true"`
}
