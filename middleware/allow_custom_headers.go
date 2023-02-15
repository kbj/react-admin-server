// Package middleware 全局中间件添加`Access-Control-Expose-Headers`的Header以便前端能读取这个Header所指定的Header内容
// @author: kbj
// @date: 2023/2/12
package middleware

import "github.com/gofiber/fiber/v2"

func AddCustomHeaders() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()

		// 以便前端能读取
		ctx.Set(fiber.HeaderAccessControlExposeHeaders, fiber.HeaderAuthorization)

		return err
	}
}
