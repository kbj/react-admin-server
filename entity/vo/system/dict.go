// Package system 字典Vo
// @author: kbj
// @date: 2023/3/1
package system

type DictSearch struct {
	DictName string `json:"dictName,omitempty"` // 字典名称
	DictType string `json:"dictType,omitempty"` // 字典类型名
	Enabled  *bool  `json:"enabled,omitempty"`  //是否启用
}
