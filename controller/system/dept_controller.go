// Package system 部门Controller
// @author: kbj
// @date: 2023/4/1
package system

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/entity/domain"
	"react-admin-server/entity/vo"
	"react-admin-server/entity/vo/system"
	"react-admin-server/service"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
)

type DeptController struct {
}

// List 列表
func (*DeptController) List(ctx *fiber.Ctx) error {
	var param domain.Dept
	if err := ctx.QueryParser(&param); err != nil {
		return err
	}
	return service.DeptService.List(ctx, &param)
}

// GetInfo 查询
func (*DeptController) GetInfo(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)
	return service.DeptService.GetInfo(ctx, id)
}

// Add 新增
func (*DeptController) Add(ctx *fiber.Ctx) error {
	var param system.DeptForm
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	}
	return service.DeptService.Add(ctx, &param)
}

// Edit 编辑
func (*DeptController) Edit(ctx *fiber.Ctx) error {
	var param system.DeptForm
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	} else if param.ID < 1 {
		return r.Fail(ctx, "ID不能为空")
	}
	return service.DeptService.Edit(ctx, &param)
}

// Delete 删除
func (*DeptController) Delete(ctx *fiber.Ctx) error {
	var ids vo.Ids
	_ = ctx.ParamsParser(&ids)
	return service.DeptService.Delete(ctx, &ids.IDs)
}
