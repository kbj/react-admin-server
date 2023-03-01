// Package system 字典Service
// @author: kbj
// @date: 2023/3/1
package system

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"react-admin-server/entity/domain"
	"react-admin-server/entity/vo/system"
	"react-admin-server/global/g"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
)

type DictService struct {
}

// List 字典列表
func (*DictService) List(ctx *fiber.Ctx, param *system.DictSearch) error {
	db := g.DbClient.Model(&domain.Dict{})

	if param.DictName != "" {
		db.Where("dict_name like ?", "%"+param.DictName+"%")
	}
	if param.DictType != "" {
		db.Where("dict_type like ?", "%"+param.DictType+"%")
	}
	if param.Enabled != nil {
		db.Where("enabled = ?", param.Enabled)
	}
	list, err := tool.SelectPageList[domain.Dict](ctx, db)
	if err != nil {
		return err
	}
	return r.Ok(ctx, r.Data(list))
}

// GetInfo 查询字典类型
func (*DictService) GetInfo(ctx *fiber.Ctx, id int) error {
	var user domain.Dict
	err := g.DbClient.First(&user, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return r.Fail(ctx, "参数有误")
	}
	return r.Ok(ctx, r.Data(&user))
}
