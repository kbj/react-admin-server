package home

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/entity/vo/home"
	"react-admin-server/service"
)

type ElectricityController struct {
}

// AddMonth 月电量统计新增
func (*ElectricityController) AddMonth(ctx *fiber.Ctx) error {
	var param []home.ElectricityMonth
	_ = ctx.BodyParser(&param)

	return service.ElectricityService.AddMonth(ctx, &param)
}

// AddDay 日电量统计新增
func (*ElectricityController) AddDay(ctx *fiber.Ctx) error {
	var param []home.ElectricityDay
	_ = ctx.BodyParser(&param)

	return service.ElectricityService.AddDay(ctx, &param)
}
