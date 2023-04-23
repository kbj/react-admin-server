package domain

type RoleMenu struct {
	RoleId uint `gorm:"comment:角色ID;index" json:"roleId"`
	MenuId uint `gorm:"comment:菜单ID;index" json:"menuId"`
}
