// Package system
// @author: kbj
// @date: 2023/2/9
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

type UserService struct {
}

// List 用户列表
func (*UserService) List(ctx *fiber.Ctx, param *system.UserSearch) error {
	db := g.DbClient.Model(&domain.User{})

	// 查询条件
	if param.Username != "" {
		db.Where("username like ?", "%"+param.Username+"%")
	}
	if param.NickName != "" {
		db.Where("nick_name like ?", "%"+param.NickName+"%")
	}
	if param.Mobile != "" {
		db.Where("mobile = ?", param.Mobile)
	}
	if param.Gender != "" {
		db.Where("gender = ?", param.Gender)
	}
	page, err := tool.SelectPageList[domain.User](ctx, db)
	if err != nil {
		return err
	}

	return r.Ok(ctx, r.Data(page))
}

// GetInfo 用户信息
func (*UserService) GetInfo(ctx *fiber.Ctx, id int) error {
	var user domain.User
	err := g.DbClient.First(&user, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return r.Fail(ctx, "参数有误")
	}
	return r.Ok(ctx, r.Data(&user))
}

// Add 新增用户
func (*UserService) Add(ctx *fiber.Ctx, param *system.UserRequest) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		// 判断用户名是否重复
		var newUser domain.User
		if err := tx.Where("username = ?", param.Username).First(&newUser).Error; err == nil || !errors.Is(err, gorm.ErrRecordNotFound) || newUser.ID > 0 {
			// 有记录或者其他错误
			return consts.NewServiceError("用户名已存在")
		}

		// 保存新用户
		newUser.Username = param.Username
		newUser.Password = tool.Md5Encode(newUser.Username+param.Password, 512)
		newUser.DeptId = param.DeptId
		newUser.Mobile = param.Mobile
		newUser.Gender = param.Gender
		newUser.NickName = param.NickName
		newUser.CreateBy = g.LoginUser.UserId(ctx)
		newUser.CreateAt = time.Now().UnixMilli()
		newUser.UpdateBy = g.LoginUser.UserId(ctx)
		newUser.UpdateAt = time.Now().UnixMilli()
		db := tx.Save(&newUser)
		if db.Error != nil || db.RowsAffected < 1 {
			_ = tool.LogDbError(db.Error)
			return consts.NewServiceError("保存失败")
		}

		// 保存角色
		if param.Roles != nil && len(*param.Roles) > 0 {
			for _, roleId := range *param.Roles {
				if err := tx.Model(&newUser).Association("Roles").Append(&domain.Role{Common: domain.Common{ID: roleId}}); err != nil {
					return err
				}
			}
		}

		// 提交事务
		return r.Ok(ctx, r.Msg("保存成功"))
	})
}

// Edit 修改用户
func (*UserService) Edit(ctx *fiber.Ctx, param *system.UserRequest) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		var user domain.User
		err := tx.First(&user, param.ID).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			return r.Fail(ctx, "参数有误")
		}

		// 保存编辑
		user.DeptId = param.DeptId
		user.Mobile = param.Mobile
		user.Gender = param.Gender
		user.NickName = param.NickName
		user.UpdateBy = g.LoginUser.UserId(ctx)
		user.UpdateAt = time.Now().UnixMilli()
		db := tx.Omit("password").Save(&user)
		if db.Error != nil || db.RowsAffected < 1 {
			_ = tool.LogDbError(db.Error)
			return consts.NewServiceError("保存失败")
		}

		// 保存角色
		if err = tx.Model(&user).Association("Roles").Clear(); err != nil {
			return err
		}
		for _, roleId := range *param.Roles {
			if err = tx.Model(&user).Association("Roles").Append(&domain.Role{Common: domain.Common{ID: roleId}}); err != nil {
				return err
			}
		}
		return r.Ok(ctx, r.Msg("保存成功"))
	})
}

// Delete 删除用户
func (*UserService) Delete(ctx *fiber.Ctx, ids *[]uint) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&domain.User{}, ids).Error; err != nil {
			return err
		}
		return r.Ok(ctx, r.Msg("删除成功"))
	})
}
