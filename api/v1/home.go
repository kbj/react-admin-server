package v1

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/controller"
)

type HomeApi struct {
	router fiber.Router
}

func (api *HomeApi) Init() {
	api.initWaterApi() // 初始化水费相关API
}

func (api *HomeApi) initWaterApi() {
	route := api.router.Group("water")
	route.Post("/", controller.WaterController.Add)
}
