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
	api.initUserApi() // 初始化用户模块
	api.initDictApi() // 初始化字典模块
}

func (api *SystemApi) initUserApi() {
	route := api.router.Group("user")
	route.Get("/list", controller.UserController.List)
	route.Get("/:id", controller.UserController.GetInfo)
	route.Post("/", controller.UserController.Add)
	route.Put("/", controller.UserController.Edit)
	route.Delete("/:ids", controller.UserController.Delete)
}

func (api *SystemApi) initDictApi() {
	// 字典管理
	route := api.router.Group("dict")
	route.Get("/list", controller.DictController.List)
	route.Get("/:id", controller.DictController.GetInfo)
	route.Post("/", controller.DictController.Add)
	route.Put("/", controller.DictController.Edit)
	route.Delete("/:ids", controller.DictController.Delete)

	// 字典值管理
	dataRoute := route.Group("data")
	dataRoute.Get("/list", controller.DictController.DataList)
	dataRoute.Get("/:id", controller.DictController.GetDataInfo)
	dataRoute.Post("/", controller.DictController.DataAdd)
	dataRoute.Put("/", controller.DictController.DataEdit)
	dataRoute.Delete("/:ids", controller.DictController.DataDelete)
	dataRoute.Get("/type/:dictType", controller.DictController.GetType)
}
