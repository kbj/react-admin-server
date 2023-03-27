// Package domain 字典类型表
// @author: kbj
// @date: 2023/2/26
package domain

type Dict struct {
	Common
	DictName string `json:"dictName,omitempty" gorm:"size:100;comment:字典名称;index;not null"`
	DictType string `json:"dictType,omitempty" gorm:"size:100;comment:字典类型名;index;not null"`
	Enabled  string `json:"enabled" gorm:"size:1;comment:是否启用;not null;default:1"`
}

type DictData struct {
	Common
	DictType  string `json:"dictType,omitempty" gorm:"size:100;comment:字典类型;index;not null"`
	DictSort  int    `json:"dictSort,omitempty" gorm:"size:10;comment:排序;not null;default:0"`
	DictLabel string `json:"dictLabel,omitempty" gorm:"size:100;comment:字典标签;not null"`
	DictValue string `json:"dictValue,omitempty" gorm:"size:100;comment:字典键值;not null"`
	TagType   string `json:"tagType,omitempty" gorm:"size:100;comment:标签类型"`
	Enabled   string `json:"enabled,omitempty" gorm:"size:1;comment:是否启用;not null;default:1"`
}
