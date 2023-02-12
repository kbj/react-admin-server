// Package g	登录用户信息
// @author: kbj
// @date: 2023/2/7
package g

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"react-admin-server/entity/domain"
)

var LoginUser loginUser

type loginUser struct {
}

// UserId 当前登录用户ID
func (*loginUser) UserId(ctx *fiber.Ctx) uint {
	return uint(ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user"].(map[string]any)["id"].(float64))
}

// User 当前登录用户信息
func (l *loginUser) User(ctx *fiber.Ctx) *domain.User {
	var user domain.User
	DbClient.First(&user, l.UserId(ctx))
	return &user
}
