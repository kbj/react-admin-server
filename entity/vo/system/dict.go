// Package system 字典Vo
// @author: kbj
// @date: 2023/3/1
package system

type DictSearch struct {
	DictName string `json:"dictName,omitempty"` // 字典名称
	DictType string `json:"dictType,omitempty"` // 字典类型名
	Enabled  *bool  `json:"enabled,omitempty"`  //是否启用
}

type DictForm struct {
	ID       uint   `json:"id" comment:"ID主键"`
	DictName string `json:"dictName,omitempty" validate:"required,min=1,max=100" comment:"字典名称"`  // 字典名称
	DictType string `json:"dictType,omitempty" validate:"required,min=1,max=100" comment:"字典类型名"` // 字典类型名
	Enabled  *bool  `json:"enabled,omitempty" validate:"required" comment:"是否启用"`                 //是否启用
}
