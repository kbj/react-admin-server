// Package system
// @author: kbj
// @date: 2023/4/8
package system

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"react-admin-server/entity/domain"
	"react-admin-server/global/g"
	"react-admin-server/global/types"
	"react-admin-server/tool/r"
	"time"
)

type MenuService struct {
}

// List 列表
func (*MenuService) List(ctx *fiber.Ctx, param *domain.Menu) error {
	db := g.DbClient.Model(&domain.Menu{})
	if param.MenuName != "" {
		db.Where("menu_name like ?", "%"+param.MenuName+"%")
	}
	if param.Enabled != "" {
		db.Where("enabled = ?", param.Enabled)
	}
	var list []domain.Menu
	db.Find(&list)
	return r.Ok(ctx, r.Data(&list))
}

// GetInfo 查询
func (*MenuService) GetInfo(ctx *fiber.Ctx, id int) error {
	var entity domain.Menu
	err := g.DbClient.First(&entity, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return r.Fail(ctx, "参数有误")
	}
	return r.Ok(ctx, r.Data(&entity))
}

// Add 新增
func (*MenuService) Add(ctx *fiber.Ctx, param *domain.Menu) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		param.ID = 0
		param.CreateBy = g.LoginUser.UserId(ctx)
		param.CreateAt = time.Now().UnixMilli()
		if err := tx.Save(param).Error; err != nil {
			return err
		}
		return r.Ok(ctx, r.Msg("保存成功"))
	})
}

// Edit 编辑
func (*MenuService) Edit(ctx *fiber.Ctx, param *domain.Menu) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		var entity domain.Menu
		if err := tx.First(&entity, param.ID).Error; err != nil {
			return err
		}
		param.ID = entity.ID
		param.CreateBy = entity.CreateBy
		param.CreateAt = entity.CreateAt
		param.UpdateBy = g.LoginUser.UserId(ctx)
		param.UpdateAt = time.Now().UnixMilli()
		if param.MenuType == types.Button {
			param.Enabled = entity.Enabled
		}
		if err := tx.Save(param).Error; err != nil {
			return err
		}
		return r.Ok(ctx, r.Msg("保存成功"))
	})
}

// Delete 删除
func (*MenuService) Delete(ctx *fiber.Ctx, param *[]int64) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		sql := `
			with recursive tmp as (
				select * from t_menu where id in ?
				union ALL
				select t1.* from t_menu t1, tmp where t1.parent_id = tmp.id and t1.delete_at = 0
			)
			update t_menu set delete_at = ?, delete_by = ? where id in (select id from tmp)
		`
		db := tx.Exec(sql, *param, time.Now().UnixMilli(), g.LoginUser.UserId(ctx))
		if db.RowsAffected > 0 {
			return r.Ok(ctx, r.Msg("删除成功"))
		}
		return db.Error
	})
}
