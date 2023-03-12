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
		db.Where("dict_type = ?", param.DictType)
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
	var dict domain.Dict
	err := g.DbClient.First(&dict, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return r.Fail(ctx, "参数有误")
	}
	return r.Ok(ctx, r.Data(&dict))
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
		for _, id := range *param {
			var entity domain.Dict
			tx.First(&entity, id)

			// 删除字典值
			if err := tx.Unscoped().Where("dict_type = ?", entity.DictType).Delete(&domain.DictData{}).Error; err != nil {
				return err
			}

			// 删除字典
			tx.Delete(&entity)
		}
		return r.Ok(ctx, r.Msg("删除成功"))
	})
}

// DataList 字典值列表
func (*DictService) DataList(ctx *fiber.Ctx, param *system.DictDataSearch) error {
	db := g.DbClient.Model(&domain.DictData{}).Where("dict_type = ?", param.DictType)

	if param.DictLabel != "" {
		db.Where("dict_label like ?", "%"+param.DictLabel+"%")
	}
	if param.Enabled != nil {
		db.Where("enabled = ?", *param.Enabled)
	}

	list, err := tool.SelectPageList[domain.DictData](ctx, db)
	if err != nil {
		return err
	}
	return r.Ok(ctx, r.Data(list))
}

// GetDataInfo 字典值详情
func (*DictService) GetDataInfo(ctx *fiber.Ctx, id int) error {
	var data domain.DictData
	if err := g.DbClient.Find(&data, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r.Fail(ctx, "参数有误")
		}
		return err
	}
	return r.Ok(ctx, r.Data(&data))
}

// AddData 增加字典值
func (*DictService) AddData(ctx *fiber.Ctx, param *system.DictDataForm) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		// 检查是否重复
		var count int64
		if err := tx.Model(&domain.DictData{}).Where("dict_type = ? and dict_value = ?", param.DictType, param.DictValue).Count(&count).Error; err != nil {
			return err
		} else if count > 0 {
			return consts.NewServiceError("该字典值已存在")
		}

		entity := domain.DictData{
			Common: domain.Common{
				CreateAt: time.Now().UnixMilli(),
				CreateBy: g.LoginUser.UserId(ctx),
			},
			DictType:  param.DictType,
			DictSort:  param.DictSort,
			DictLabel: param.DictLabel,
			DictValue: param.DictValue,
			TagType:   param.TagType,
			Enabled:   param.Enabled,
		}
		db := tx.Save(&entity)
		if db.Error != nil || db.RowsAffected < 1 {
			_ = tool.LogDbError(db.Error)
			return consts.NewServiceError("保存失败")
		}
		return r.Ok(ctx)
	})
}

// DataEdit 编辑字典值
func (*DictService) DataEdit(ctx *fiber.Ctx, param *system.DictDataForm) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		var dict domain.DictData
		if err := tx.First(&dict, param.ID).Error; err != nil {
			return err
		}
		dict.UpdateBy = g.LoginUser.UserId(ctx)
		dict.UpdateAt = time.Now().UnixMilli()
		dict.DictType = param.DictType
		dict.DictSort = param.DictSort
		dict.TagType = param.TagType
		dict.Enabled = param.Enabled
		dict.DictLabel = param.DictLabel
		dict.DictValue = param.DictValue
		if tx.Save(&dict).RowsAffected < 1 {
			return consts.NewServiceError("更新失败")
		}
		return r.Ok(ctx, r.Msg("保存成功"))
	})
}

// DataDelete 删除字典值
func (*DictService) DataDelete(ctx *fiber.Ctx, param *[]uint) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		// 删除字典
		var count int64
		db := tx.Unscoped().Delete(&domain.DictData{}, param)
		db.Count(&count)
		if err := db.Error; err != nil {
			return err
		} else if count < 1 {
			return consts.NewServiceError("删除失败")
		}
		return r.Ok(ctx, r.Msg("删除成功"))
	})
}

// GetType 通用查询字典类型的列表信息
func (*DictService) GetType(ctx *fiber.Ctx, dictType string) error {
	var dictList []domain.DictData
	if err := g.DbClient.Where("dict_type = ? and enabled = true", dictType).Order("dict_sort, id").Find(&dictList).Error; err != nil {
		return err
	}
	return r.Ok(ctx, r.Data(&dictList))
}
