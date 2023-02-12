// Package cmd	校验器的初始化
// @author: kbj
// @date: 2023/2/3
package cmd

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"react-admin-server/global/g"
	"reflect"
)

func initValidator() (*ut.Translator, *validator.Validate) {
	//设置支持语言
	chinese := zh.New()
	english := en.New()
	//设置国际化翻译器
	uni := ut.New(chinese, chinese, english)
	//设置验证器
	val := validator.New()
	//根据参数取翻译器实例
	trans, _ := uni.GetTranslator("zh")
	// 注册中文为默认语言
	if err := zhTranslations.RegisterDefaultTranslations(val, trans); err != nil {
		g.Logger.Error("校验器注册中文语言失败", zap.Error(err))
	}
	//使用fld.Tag.Get("comment")注册一个获取tag的自定义方法
	val.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})
	return &trans, val
}
