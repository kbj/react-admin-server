// Package system 部门Service
// @author: kbj
// @date: 2023/4/1
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
	"strconv"
	"time"
)

type DeptService struct {
}

// List 部门列表
func (*DeptService) List(ctx *fiber.Ctx, param *domain.Dept) error {
	db := g.DbClient.Model(&domain.Dept{})
	if param.DeptName != "" {
		db.Where("dept_name like ?", "%"+param.DeptName+"%")
	}
	if param.Enabled != "" {
		db.Where("enabled = ?", param.Enabled)
	}

	var list []domain.Dept
	db.Find(&list)
	return r.Ok(ctx, r.Data(&list))
}

// GetInfo 查询
func (*DeptService) GetInfo(ctx *fiber.Ctx, id int) error {
	var entity domain.Dept
	err := g.DbClient.First(&entity, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return r.Fail(ctx, "参数有误")
	}
	return r.Ok(ctx, r.Data(&entity))
}

// Add 新增
func (*DeptService) Add(ctx *fiber.Ctx, param *system.DeptForm) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		// 查询父部门
		var parent domain.Dept
		if err := tx.First(&parent, param.ParentId).Error; err != nil {
			return err
		}

		entity := domain.Dept{
			Common: domain.Common{
				CreateBy: g.LoginUser.UserId(ctx),
				CreateAt: time.Now().UnixMilli(),
			},
			Ancestors:    parent.Ancestors + "," + strconv.Itoa(int(param.ParentId)),
			DeptName:     param.DeptName,
			ParentId:     param.ParentId,
			OrderNum:     param.OrderNum,
			LeaderUserId: param.LeaderUserId,
			Enabled:      param.Enabled,
		}
		db := tx.Save(&entity)
		if db.Error != nil || db.RowsAffected < 1 {
			_ = tool.LogDbError(db.Error)
			return consts.NewServiceError("保存失败")
		}
		return r.Ok(ctx)
	})
}

// Edit 编辑
func (*DeptService) Edit(ctx *fiber.Ctx, param *system.DeptForm) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		var entity domain.Dept
		if err := tx.First(&entity, param.ID).Error; err != nil {
			return err
		}
		var parent domain.Dept
		if err := tx.First(&parent, param.ParentId).Error; err != nil {
			return err
		}

		entity.Ancestors = parent.Ancestors + "," + strconv.Itoa(int(param.ParentId))
		entity.DeptName = param.DeptName
		entity.ParentId = param.ParentId
		entity.OrderNum = param.OrderNum
		entity.LeaderUserId = param.LeaderUserId
		entity.Enabled = param.Enabled
		entity.UpdateBy = g.LoginUser.UserId(ctx)
		entity.UpdateAt = time.Now().UnixMilli()
		if tx.Save(&entity).RowsAffected < 1 {
			return consts.NewServiceError("更新失败")
		}
		return r.Ok(ctx, r.Msg("保存成功"))
	})
}

// Delete 删除
func (*DeptService) Delete(ctx *fiber.Ctx, param *[]uint) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		sql := `
			with recursive tmp as (
				select * from t_dept where id in ?
				union ALL
				select t1.* from t_dept t1, tmp where t1.parent_id = tmp.id and t1.delete_at > 0
			)
			update t_dept set delete_at = now() and delete_by = ? where id in (select id from tmp)
		`
		db := tx.Exec(sql, param, g.LoginUser.UserId(ctx))
		if db.RowsAffected > 0 {
			return r.Ok(ctx, r.Msg("删除成功"))
		}
		return db.Error
	})
}
