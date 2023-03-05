// Package system 用户管理模块
// @author: kbj
// @date: 2023/2/9
package system

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/entity/vo"
	"react-admin-server/entity/vo/system"
	"react-admin-server/service"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
)

type UserController struct {
}

// List 用户列表
func (*UserController) List(ctx *fiber.Ctx) error {
	var param system.UserSearch
	if err := ctx.QueryParser(&param); err != nil {
		return err
	}
	return service.UserService.List(ctx, &param)
}

// GetInfo 用户信息
func (*UserController) GetInfo(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)
	return service.UserService.GetInfo(ctx, id)
}

// Add 新增用户
func (*UserController) Add(ctx *fiber.Ctx) error {
	var param system.UserRequest
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	}
	return service.UserService.Add(ctx, &param)
}

// Edit 修改用户
func (*UserController) Edit(ctx *fiber.Ctx) error {
	var param system.UserRequest
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	} else if param.ID == 0 {
		return r.Fail(ctx, "ID为空")
	}
	return service.UserService.Edit(ctx, &param)
}

// Delete 删除用户
func (*UserController) Delete(ctx *fiber.Ctx) error {
	var params vo.Ids
	_ = ctx.ParamsParser(&params)
	return service.UserService.Delete(ctx, &params.IDs)
}
