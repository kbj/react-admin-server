// Package domain 字典类型表
// @author: kbj
// @date: 2023/2/26
package domain

type Dict struct {
	Common
	DictName string `json:"dictName,omitempty" gorm:"size:100;comment:字典名称;index;not null"`
	DictType string `json:"dictType,omitempty" gorm:"size:100;comment:字典类型名;index;not null"`
	Enabled  bool   `json:"enabled" gorm:"comment:是否启用;not null;default:true"`
}

type DictData struct {
	Common
	Dict      Dict   `gorm:"foreignKey:DictType;type:string"`
	DictType  string `json:"dictType,omitempty" gorm:"size:100;comment:字典类型名;index;not null"`
	DictSort  int    `json:"dictSort,omitempty" gorm:"size:10;comment:排序;not null;default:0"`
	DictLabel string `json:"dictLabel,omitempty" gorm:"size:100;comment:字典名称;not null"`
	DictValue string `json:"dictValue,omitempty" gorm:"size:100;comment:字典值;not null"`
	Enabled   bool   `json:"enabled,omitempty" gorm:"comment:是否启用;not null;default:true"`
}
