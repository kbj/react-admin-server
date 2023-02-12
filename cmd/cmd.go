// Package cmd 初始化层逻辑
// @author: kbj
// @date: 2023/1/30
package cmd

import (
	"react-admin-server/global/g"
)

// InitStart 初始化方法
func InitStart() {
	g.V = initViper()                                    // 初始化Viper
	g.Logger = initZap()                                 // 初始化Zap
	g.DbClient = initGorm()                              // 初始化数据库
	g.ValidatorTranslator, g.Validator = initValidator() // 初始化参数校验库
	g.FiberApp = initFiber()                             // 初始化Fiber的App对象

	boot() // 启动Fiber
}
