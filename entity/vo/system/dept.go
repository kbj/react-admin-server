// Package system 部门Vo
// @author: kbj
// @date: 2023/4/1
package system

type DeptForm struct {
	ID           uint   `json:"id,omitempty" comment:"主键"`
	DeptName     string `json:"deptName,omitempty" validate:"required,min=1,max=100" comment:"部门名称"`
	ParentId     uint   `json:"parentId,omitempty" validate:"required" comment:"父ID"`
	OrderNum     int    `json:"orderNum,omitempty" comment:"排序"`
	LeaderUserId uint   `json:"leaderUserId,omitempty" comment:"负责人ID"`
	Enabled      string `json:"enabled,omitempty" comment:"是否启用"`
}
