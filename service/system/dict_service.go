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
	"react-admin-server/global/consts"
	"react-admin-server/global/g"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
	"time"
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

// Add 新增字典类型
func (*DictService) Add(ctx *fiber.Ctx, param *system.DictForm) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		// 查询是否有重复
		var count int64
		tx.Model(&domain.Dict{}).Where("dict_type = ?", param.DictType).Count(&count)
		if count > 0 {
			return consts.NewServiceError("字典类型名称已重复")
		}

		// 保存记录
		dict := domain.Dict{
			DictName: param.DictName,
			DictType: param.DictType,
			Enabled:  *param.Enabled,
		}
		dict.CreateBy = g.LoginUser.UserId(ctx)
		dict.CreateAt = time.Now().UnixMilli()
		db := tx.Save(&dict)
		if db.Error != nil || db.RowsAffected < 1 {
			_ = tool.LogDbError(db.Error)
			return consts.NewServiceError("保存失败")
		}
		return r.Ok(ctx)
	})
}

// Edit 编辑
func (*DictService) Edit(ctx *fiber.Ctx, param *system.DictForm) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		var dict domain.Dict
		if err := tx.First(&dict, param.ID).Error; err != nil {
			return err
		} else if dict.DictType != param.DictType {
			// 查询类型是否重复
			var count int64
			if err = tx.Model(domain.Dict{}).Where("dict_type = ?", param.DictType).Count(&count).Error; err != nil {
				return err
			} else if count > 0 {
				return consts.NewServiceError("字典类型已重复")
			}
		}
		dict.UpdateBy = g.LoginUser.UserId(ctx)
		dict.UpdateAt = time.Now().UnixMilli()
		dict.DictType = param.DictType
		dict.DictName = param.DictName
		dict.Enabled = *param.Enabled
		if tx.Save(&dict).RowsAffected < 1 {
			return consts.NewServiceError("更新失败")
		}
		return r.Ok(ctx, r.Msg("保存成功"))
	})
}

// Delete 删除
func (*DictService) Delete(ctx *fiber.Ctx, param *[]uint) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&domain.Dict{}, param).Error; err != nil {
			return err
		}
		return r.Ok(ctx, r.Msg("删除成功"))
	})
}
