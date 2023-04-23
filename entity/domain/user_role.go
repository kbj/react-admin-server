package domain

type UserRole struct {
	UserId uint `gorm:"comment:用户ID;index" json:"userId"`
	RoleId uint `gorm:"comment:角色ID;index" json:"roleId"`
}
