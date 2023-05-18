// Package middleware JWT中间件
// @author: kbj
// @date: 2023/2/5
package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"react-admin-server/global/g"
	"react-admin-server/service"
	"react-admin-server/tool/r"
	"time"
)

// InitJwt 初始化JWT
func InitJwt() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(g.Env.Jwt.Key),
		SigningMethod:  "HS512",
		ErrorHandler:   jwtCheckError,
		SuccessHandler: jwtSuccessHandler,
		Filter:         jwtByPass,
	})
}

// Jwt校验失败处理逻辑
func jwtCheckError(ctx *fiber.Ctx, err error) error {
	return r.Forbidden(ctx)
}

// Jwt校验通过的钩子
func jwtSuccessHandler(ctx *fiber.Ctx) error {
	err := ctx.Next()

	// 检查token是否将要过期
	users := ctx.Locals("user").(*jwt.Token)
	claims := users.Claims.(jwt.MapClaims)
	exp := int64(claims["exp"].(float64))
	if time.Now().Add(time.Minute*3).Unix()-exp >= 0 {
		// 有效期在3分钟以内，自动续期
		user := claims["user"].(map[string]any)
		token := service.LoginService.RefreshToken(user["id"].(uint))
		if token != "" {
			ctx.Set(fiber.HeaderAuthorization, token)
		}
	}
	return err
}

// 绕过Jwt校验的逻辑
func jwtByPass(ctx *fiber.Ctx) bool {
	path := ctx.Path()
	for _, rule := range g.Env.Jwt.ByPass {
		if fiber.RoutePatternMatch(path, rule) {
			return true
		}
	}
	return false
}
