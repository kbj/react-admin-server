// Package system 角色服务
// @author: kbj
// @date: 2023/4/12
package system

import (
	"errors"
	"fmt"
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

type RoleService struct {
}

// List 列表
func (*RoleService) List(ctx *fiber.Ctx, param *system.RoleSearch) error {
	db := g.DbClient.Model(&domain.Role{})

	if param.RoleName != nil {
		db.Where("role_name like ?", "%"+*param.RoleName+"%")
	}
	if param.RoleKey != nil {
		db.Where("role_key like ?", "%"+*param.RoleKey+"%")
	}
	if param.Enabled != nil {
		db.Where("enabled is ?", *param.Enabled)
	}

	list, err := tool.SelectPageList[domain.Role](ctx, db)
	if err != nil {
		return err
	}
	return r.Ok(ctx, r.Data(list))
}

// GetInfo 查询
func (*RoleService) GetInfo(ctx *fiber.Ctx, id int) error {
	var entity system.RoleForm
	err := g.DbClient.Table("t_role").First(&entity, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return r.Fail(ctx, "参数有误")
	}

	// 查询关联菜单ID
	var menus []uint
	g.DbClient.Raw("select menu_id from t_role_menu where role_id = ?", entity.ID).Scan(&menus)
	entity.Menus = &menus
	return r.Ok(ctx, r.Data(&entity))
}

// Add 新增
func (*RoleService) Add(ctx *fiber.Ctx, param *system.RoleForm) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.Model(domain.Role{}).Where("role_key = ?", param.RoleKey).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return consts.NewServiceError("权限字符已重复，请勿重复创建")
		}

		// 构建实体
		entity := domain.Role{
			Common: domain.Common{
				CreateAt: time.Now().UnixMilli(),
				CreateBy: g.LoginUser.UserId(ctx),
			},
			RoleName: param.RoleName,
			RoleKey:  param.RoleKey,
			Enabled:  param.Enabled,
			OrderNum: param.OrderNum,
		}
		if err := tx.Create(&entity).Error; err != nil {
			return err
		}

		// 菜单
		if param.Menus != nil && len(*param.Menus) > 0 {
			var menus domain.RoleMenu
			menus.RoleId = entity.ID

			for _, id := range *param.Menus {
				menus.MenuId = id
				if err := tx.Create(&menus).Error; err != nil {
					return err
				}
			}
		}

		return r.Ok(ctx, r.Msg("保存成功"))
	})
}

// Edit 编辑
func (*RoleService) Edit(ctx *fiber.Ctx, param *system.RoleForm) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		var entity domain.Role
		if err := tx.First(&entity, param.ID).Error; err != nil {
			return err
		}

		// 清空菜单关联
		if err := tx.Where("role_id = ?", param.ID).Delete(&domain.RoleMenu{}).Error; err != nil {
			return err
		}

		// 更新角色
		entity.UpdateBy = g.LoginUser.UserId(ctx)
		entity.UpdateAt = time.Now().UnixMilli()
		entity.RoleName = param.RoleName
		entity.Enabled = param.Enabled
		entity.OrderNum = param.OrderNum
		if err := tx.Save(&entity).Error; err != nil {
			return err
		}
		// 更新菜单
		if param.Menus != nil && len(*param.Menus) > 0 {
			var menus domain.RoleMenu
			menus.RoleId = entity.ID

			for _, id := range *param.Menus {
				menus.MenuId = id
				if err := tx.Create(&menus).Error; err != nil {
					return err
				}
			}
		}
		return r.Ok(ctx, r.Msg("保存成功"))
	})
}

// Delete 删除
func (*RoleService) Delete(ctx *fiber.Ctx, param *[]int64) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		for _, i := range *param {
			// 需要检查角色是否有被用户关联
			var userNum int64
			tx.Model(domain.UserRole{}).Where("role_id=?", i).Count(&userNum)
			if userNum > 0 {
				var entity domain.Role
				if err := tool.LogDbError(tx.First(&entity, i).Error); err != nil {
					return err
				}
				return consts.NewServiceError(fmt.Sprintf("%s 已分配,不能删除", *entity.RoleName))
			}

			// 删除本体与菜单关联
			if err := tool.LogDbError(tx.Where("role_id = ?", i).Delete(&domain.RoleMenu{}).Error); err != nil {
				return consts.NewServiceError("删除失败")
			}
			if err := tool.LogDbError(tx.Delete(&domain.RoleMenu{}, i).Error); err != nil {
				return consts.NewServiceError("删除失败")
			}
		}
		return r.Ok(ctx, r.Msg("删除成功"))
	})
}
