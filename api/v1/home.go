package v1

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/controller"
)

type HomeApi struct {
	router fiber.Router
}

func (api *HomeApi) Init() {
	api.initWaterApi()       // 初始化水费相关API
	api.initElectricityApi() // 初始化电费相关API
}

func (api *HomeApi) initWaterApi() {
	route := api.router.Group("water")
	route.Post("/", controller.WaterController.Add)
}

func (api *HomeApi) initElectricityApi() {
	route := api.router.Group("electricity")
	route.Post("/month", controller.ElectricityController.AddMonth)
	route.Post("/day", controller.ElectricityController.AddDay)
}
