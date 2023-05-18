// Package v1 系统通用接口
// @author: kbj
// @date: 2023/2/9
package v1

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/controller"
	"react-admin-server/global/g"
)

type CommonApi struct {
	router *fiber.Router
}

func (api *CommonApi) Init() {
	api.initLoginApi() // 用户相关模块
	api.initFileApi()  // 文件相关模块
}

func (api *CommonApi) initLoginApi() {
	route := *api.router
	route.Post("/login", controller.LoginController.Login)       // 登录接口
	route.Get("/user-info", controller.LoginController.UserInfo) // 查询登录用户信息
	route.Get("/menus", controller.LoginController.Menus)        // 查询登录用户菜单信息
	route.Get("/roles", controller.LoginController.RolesList)    // 查询系统所有角色信息
}

func (api *CommonApi) initFileApi() {
	route := (*api.router).Group("common/file")
	route.Post("/upload", controller.FileController.UploadFile) // 上传文件
	route.Static("/static", g.Env.System.UploadPath)            // 静态目录注册
}
