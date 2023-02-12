// Package cmd Fiber框架的相关初始化
// @author: kbj
// @date: 2023/1/31
package cmd

import (
	"fmt"
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"os"
	"os/signal"
	v1 "react-admin-server/api/v1"
	"react-admin-server/global/consts"
	"react-admin-server/global/g"
	"react-admin-server/tool/r"
	"syscall"
)

// 初始化Fiber
func initFiber() *fiber.App {
	// 初始化Fiber对象
	app := fiber.New(fiberCreateConfig())

	// 注册全局中间件
	fiberRegisterGlobalMiddleware(app)

	// 注册路由
	v1.Init(app)

	return app
}

// 启动Fiber
func boot() {
	fiberGracefullyShutDown(g.FiberApp) // 优雅关机
	if err := g.FiberApp.Listen(fmt.Sprintf("%s:%s", g.Env.System.Host, g.Env.System.Port)); err != nil {
		g.Logger.Error("Fiber启动失败", zap.Error(err))
		os.Exit(1)
	}
}

// 初始化Fiber的配置
func fiberCreateConfig() fiber.Config {
	return fiber.Config{
		AppName:      g.Env.System.ProjectName,
		BodyLimit:    g.Env.System.BodyLimit * 1024 * 1024,
		JSONEncoder:  jsoniter.Marshal,   // 使用jsoniter库
		JSONDecoder:  jsoniter.Unmarshal, // 使用jsoniter库
		ErrorHandler: fiberErrorHandler,
	}
}

// 注册全局中间件
func fiberRegisterGlobalMiddleware(app *fiber.App) {
	app.Use(
		fiberzap.New(fiberzap.Config{Logger: g.Logger}), // zap日志中间件
		recover.New(), // 恢复panic
		cors.New(),    // 跨域配置
	)
}

// 业务上全局错误处理
func fiberErrorHandler(ctx *fiber.Ctx, err error) error {
	httpCode := fiber.StatusInternalServerError

	if _, ok := err.(consts.ServiceError); ok {
		// 自定义业务异常，不需要处理
		httpCode = fiber.StatusOK
	} else if e, ok := err.(*fiber.Error); ok {
		httpCode = e.Code
		if httpCode != fiber.StatusNotFound && httpCode != fiber.StatusMethodNotAllowed && g.Logger != nil {
			// 找不到资源、请求方法不允许的错误不用打印
			g.Logger.Error("Fiber框架异常", zap.Error(err))
		}
	} else if g.Logger != nil {
		// 其他未知异常
		g.Logger.Error("未知异常", zap.Error(err))
	}

	ctx.Status(httpCode)
	return r.Fail(ctx, err.Error())
}

// 实现Fiber的优雅关机
func fiberGracefullyShutDown(app *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		_ = <-c
		g.Logger.Info("正在停止运行中...")
		cleanupTasks()
		_ = app.Shutdown()
	}()
}

// 优雅关机后业务方面需要执行的任务
func cleanupTasks() {
	// 数据库关闭连接
	if g.DbClient != nil {
		db, err := g.DbClient.DB()
		err = db.Close()
		if err != nil {
			g.Logger.Error("关闭数据库连接失败：", zap.Error(err))
		}
	}
}
