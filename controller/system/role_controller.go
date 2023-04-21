// Package system 角色控制器
// @author: kbj
// @date: 2023/4/12
package system

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/entity/vo"
	"react-admin-server/entity/vo/system"
	"react-admin-server/service"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
)

type RoleController struct {
}

// List 列表
func (*RoleController) List(ctx *fiber.Ctx) error {
	var param system.RoleSearch
	if err := ctx.QueryParser(&param); err != nil {
		return err
	}
	return service.RoleService.List(ctx, &param)
}

// GetInfo 查询
func (*RoleController) GetInfo(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)
	return service.RoleService.GetInfo(ctx, id)
}

// Add 新增
func (*RoleController) Add(ctx *fiber.Ctx) error {
	var param system.RoleForm
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	}
	return service.RoleService.Add(ctx, &param)
}

// Edit 编辑
func (*RoleController) Edit(ctx *fiber.Ctx) error {
	var param system.RoleForm
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	} else if param.ID == nil {
		return r.Fail(ctx, "ID不能为空")
	}
	return service.RoleService.Edit(ctx, &param)
}

// Delete 删除
func (*RoleController) Delete(ctx *fiber.Ctx) error {
	var ids vo.Ids
	_ = ctx.ParamsParser(&ids)
	return service.RoleService.Delete(ctx, &ids.IDs)
}
