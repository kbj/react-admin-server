// Package v1 系统模块的API
// @author: kbj
// @date: 2023/2/2
package v1

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/controller"
)

type SystemApi struct {
	router fiber.Router
}

func (api *SystemApi) Init() {
	api.initUserApi() // 初始化用户模块接口
}

func (api *SystemApi) initUserApi() {
	route := api.router.Group("user")
	route.Get("/list", controller.UserController.List)
	route.Get("/:id", controller.UserController.GetInfo)
	route.Post("/", controller.UserController.Add)
	route.Put("/", controller.UserController.Edit)
	route.Delete("/:ids", controller.UserController.Delete)
}
