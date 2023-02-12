// Package system	登录方面的Controller
// @author: kbj
// @date: 2023/2/2
package system

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/entity/vo/system"
	"react-admin-server/service"
	"react-admin-server/tool"
)

type LoginController struct {
}

// Login 登录方法
func (*LoginController) Login(ctx *fiber.Ctx) error {
	var param system.LoginRequest
	_ = ctx.BodyParser(&param)
	if err := tool.ValidateParams(&param); err != nil {
		return err
	}
	return service.LoginService.Login(ctx, &param)
}

// UserInfo 查询登录用户信息
func (*LoginController) UserInfo(ctx *fiber.Ctx) error {
	return service.LoginService.Info(ctx)
}

// Menus 用户菜单信息
func (*LoginController) Menus(ctx *fiber.Ctx) error {
	return service.LoginService.Menus(ctx)
}
