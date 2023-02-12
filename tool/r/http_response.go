// Package r	对Fiber返回数据部分的封装
// @author: kbj
// @date: 2023/2/3
package r

import (
	"github.com/gofiber/fiber/v2"
	"react-admin-server/entity/vo"
	"react-admin-server/global/consts"
)

type ResponseOptions func(response *vo.Response)

const (
	noAuthMsg   = "您暂时没有访问此资源的权限"
	notFoundMsg = "您请求的资源不存在"
)

// Msg 设置自定义返回Msg的方法
func Msg(msg string) ResponseOptions {
	return func(r *vo.Response) {
		r.Msg = msg
	}
}

// Code 设置自定义的编码值
func Code(code int) ResponseOptions {
	return func(r *vo.Response) {
		r.Code = code
	}
}

// Data 设置自定义返回Data
func Data(data any) ResponseOptions {
	return func(r *vo.Response) {
		r.Data = data
	}
}

// Ok 返回成功数据
func Ok(c *fiber.Ctx, options ...ResponseOptions) error {
	r := &vo.Response{
		Code: consts.StatusSuccess,
	}

	// 遍历可选参数调用修改
	for _, op := range options {
		op(r)
	}

	return c.JSON(r)
}

// Fail 返回失败数据
func Fail(c *fiber.Ctx, msg string) error {
	return c.JSON(&vo.Response{
		Code: consts.StatusFailure,
		Msg:  msg,
	})
}

// Forbidden 未授权访问
func Forbidden(c *fiber.Ctx, options ...ResponseOptions) error {
	r := &vo.Response{
		Code: consts.StatusFailure,
		Msg:  noAuthMsg,
	}

	// 遍历可选参数调用修改
	for _, op := range options {
		op(r)
	}

	c.Status(fiber.StatusForbidden)
	return c.JSON(r)
}

// NeedLogin 需要登录
func NeedLogin(c *fiber.Ctx) error {
	return Ok(c, Msg("登录信息已过期"), Code(consts.StatusNeedLogin))
}

// NotFound 找不到资源
func NotFound(c *fiber.Ctx) error {
	return c.JSON(&vo.Response{
		Code: consts.StatusSuccess,
		Msg:  notFoundMsg,
	})
}
