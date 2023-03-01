// Package domain 部门
// @author: kbj
// @date: 2023/2/10
package domain

type Dept struct {
	Common
	DeptName     string `json:"deptName,omitempty" gorm:"size:100;comment:部门名称;not null"`
	ParentId     uint   `json:"parentId,omitempty" gorm:"comment:父部门ID;not null;default:0"`
	OrderNum     int    `json:"orderNum,omitempty" gorm:"comment:排序;not null;default:0"`
	LeaderUserId uint   `json:"leaderUserId,omitempty" gorm:"comment:负责人ID"`
	Enabled      bool   `json:"enabled,omitempty" gorm:"comment:是否启用;not null;default:true"`
}
