// Package system 字典相关Controller
// @author: kbj
// @date: 2023/3/1
package system

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/entity/vo/system"
	"react-admin-server/service"
	"react-admin-server/tool"
)

type DictController struct {
}

// List 字典列表
func (*DictController) List(ctx *fiber.Ctx) error {
	var param system.DictSearch
	if err := ctx.QueryParser(&param); err != nil {
		return err
	}
	return service.DictService.List(ctx, &param)
}

// GetInfo 查询字典类型
func (*DictController) GetInfo(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)
	return service.DictService.GetInfo(ctx, id)
}

// Add 新增
func (*DictController) Add(ctx *fiber.Ctx) error {
	var param system.DictForm
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	}
	return service.DictService.Add(ctx, &param)
}
