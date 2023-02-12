// Package g 全局变量
// @author: kbj
// @date: 2023/1/30
package g

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"react-admin-server/cmd/env"
)

var (
	Env      env.Environment // 环境配置文件
	V        *viper.Viper    // 全局Viper对象
	Logger   *zap.Logger     // 日志对象
	DbClient *gorm.DB        // 数据库操作对象
	FiberApp *fiber.App      // Fiber App

	ValidatorTranslator *ut.Translator      // 校验器语言翻译器
	Validator           *validator.Validate // 校验器
)
