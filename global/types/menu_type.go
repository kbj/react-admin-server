// Package types 菜单类型
// @author: kbj
// @date: 2023/2/10
package types

type MenuType string

const (
	Content MenuType = "C" // 目录菜单
	Menu    MenuType = "M" // 菜单
	Button  MenuType = "B" // 按钮
)
