// Package system 字典Vo
// @author: kbj
// @date: 2023/3/1
package system

type DictSearch struct {
	DictName string `json:"dictName,omitempty"` // 字典名称
	DictType string `json:"dictType,omitempty"` // 字典类型名
	Enabled  string `json:"enabled,omitempty"`  //是否启用
}

type DictForm struct {
	ID       uint   `json:"id" comment:"ID主键"`
	DictName string `json:"dictName,omitempty" validate:"required,min=1,max=100" comment:"字典名称"`  // 字典名称
	DictType string `json:"dictType,omitempty" validate:"required,min=1,max=100" comment:"字典类型名"` // 字典类型名
	Enabled  string `json:"enabled,omitempty" validate:"required,max=1" comment:"是否启用"`           //是否启用
}

type DictDataSearch struct {
	DictType  string `json:"dictType,omitempty" validate:"required,min=1,max=100" comment:"字典类型名"` // 字典类型名
	DictLabel string `json:"dictLabel,omitempty" validate:"max=100" comment:"字典名称"`                // 字典名称
	Enabled   string `json:"enabled,omitempty"`                                                    //是否启用
}

type DictDataForm struct {
	ID        uint   `json:"id" comment:"ID主键"`
	DictType  string `json:"dictType,omitempty" validate:"required,min=1,max=100" comment:"字典类型"`
	DictSort  int    `json:"dictSort,omitempty"`
	DictLabel string `json:"dictLabel,omitempty" validate:"required,min=1,max=100" comment:"字典标签'"`
	DictValue string `json:"dictValue,omitempty" validate:"required,min=1,max=100" comment:"字典键值"`
	TagType   string `json:"tagType,omitempty"`
	Enabled   string `json:"enabled,omitempty"`
}
