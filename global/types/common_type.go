// Package types 自定义枚举类型
// @author: kbj
// @date: 2023/2/10
package types

// Gender 性别
type Gender string

const (
	Man   Gender = "M" // 男
	Woman Gender = "W" // 女
)

// MenuType 菜单类型
type MenuType string

const (
	Content MenuType = "C" // 目录菜单
	Menu    MenuType = "M" // 菜单
	Button  MenuType = "B" // 按钮
)
