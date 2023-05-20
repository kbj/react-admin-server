package home

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/entity/domain"
	"react-admin-server/service"
)

type WaterController struct {
}

// Add 新增
func (*WaterController) Add(ctx *fiber.Ctx) error {
	var param []domain.Water
	_ = ctx.BodyParser(&param)
	return service.WaterService.Add(ctx, &param)
}
