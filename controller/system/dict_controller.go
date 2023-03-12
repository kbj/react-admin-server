// Package system 字典相关Controller
// @author: kbj
// @date: 2023/3/1
package system

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/entity/vo"
	"react-admin-server/entity/vo/system"
	"react-admin-server/service"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
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

// Edit 编辑
func (*DictController) Edit(ctx *fiber.Ctx) error {
	var param system.DictForm
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	} else if param.ID < 1 {
		return r.Fail(ctx, "ID不能为空")
	}
	return service.DictService.Edit(ctx, &param)
}

// Delete 删除
func (*DictController) Delete(ctx *fiber.Ctx) error {
	var ids vo.Ids
	_ = ctx.ParamsParser(&ids)
	return service.DictService.Delete(ctx, &ids.IDs)
}

// DataList 字典值列表
func (*DictController) DataList(ctx *fiber.Ctx) error {
	var param system.DictDataSearch
	_ = ctx.QueryParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	}
	return service.DictService.DataList(ctx, &param)
}

// GetDataInfo 字典值详情
func (*DictController) GetDataInfo(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)
	return service.DictService.GetDataInfo(ctx, id)
}

// DataAdd 新增字典值
func (*DictController) DataAdd(ctx *fiber.Ctx) error {
	var param system.DictDataForm
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	}
	return service.DictService.AddData(ctx, &param)
}

// DataEdit 编辑字典值
func (*DictController) DataEdit(ctx *fiber.Ctx) error {
	var param system.DictDataForm
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	} else if param.ID < 1 {
		return r.Fail(ctx, "ID不能为空")
	}
	return service.DictService.DataEdit(ctx, &param)
}

// DataDelete 删除字典值
func (*DictController) DataDelete(ctx *fiber.Ctx) error {
	var ids vo.Ids
	_ = ctx.ParamsParser(&ids)
	return service.DictService.DataDelete(ctx, &ids.IDs)
}

// GetType 通用查询字典类型的列表信息
func (*DictController) GetType(ctx *fiber.Ctx) error {
	dictType := ctx.Params("dictType", "")
	return service.DictService.GetType(ctx, dictType)
}
