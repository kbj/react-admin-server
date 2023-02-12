// Package v1
// @author: kbj
// @date: 2023/2/2
package v1

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/middleware"
)

func Init(app *fiber.App) {
	v1Group := app.Group("v1", middleware.InitJwt())

	(&CommonApi{router: &v1Group}).Init()                // 系统通用API
	(&SystemApi{router: v1Group.Group("system")}).Init() // 系统管理
}
