// Package system 菜单
// @author: kbj
// @date: 2023/4/7
package system

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/entity/domain"
	"react-admin-server/entity/vo"
	"react-admin-server/service"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
)

type MenuController struct {
}

// List 列表
func (*MenuController) List(ctx *fiber.Ctx) error {
	var param domain.Menu
	if err := ctx.QueryParser(&param); err != nil {
		return err
	}
	return service.MenuService.List(ctx, &param)
}

// GetInfo 查询
func (*MenuController) GetInfo(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)
	return service.MenuService.GetInfo(ctx, id)
}

// Add 新增
func (*MenuController) Add(ctx *fiber.Ctx) error {
	var param domain.Menu
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	}
	return service.MenuService.Add(ctx, &param)
}

// Edit 编辑
func (*MenuController) Edit(ctx *fiber.Ctx) error {
	var param domain.Menu
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	} else if param.ID < 1 {
		return r.Fail(ctx, "ID不能为空")
	}
	return service.MenuService.Edit(ctx, &param)
}

// Delete 删除
func (*MenuController) Delete(ctx *fiber.Ctx) error {
	var ids vo.Ids
	_ = ctx.ParamsParser(&ids)
	return service.MenuService.Delete(ctx, &ids.IDs)
}
