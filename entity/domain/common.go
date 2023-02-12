// Package domain 通用结构
// @author: kbj
// @date: 2023/2/10
package domain

import "gorm.io/plugin/soft_delete"

type Common struct {
	ID       uint                  `gorm:"comment:主键;primarykey" json:"id"`
	CreateAt int64                 `gorm:"comment:创建时间;autoUpdateTime:milli;not null;<-:create" json:"createAt"`
	CreateBy uint                  `gorm:"comment:创建人;not null;<-:create" json:"createBy"`
	UpdateAt int64                 `gorm:"comment:修改时间;autoUpdateTime:milli" json:"updateAt"`
	UpdateBy uint                  `gorm:"comment:修改人" json:"updateBy"`
	DeleteAt soft_delete.DeletedAt `gorm:"comment:删除时间;index;softDelete:milli;default:0;not null" json:"deleteAt,omitempty"`
	DeleteBy uint                  `gorm:"comment:删除人" json:"deleteBy"`
}
