// Package env 环境
// @author: kbj
// @date: 2023/1/30
package env

import "react-admin-server/cmd/env/internal"

type Environment struct {
	System internal.System `json:"system" yaml:"system" mapstructure:"system"` // 系统基础配置
	Log    internal.Log    `json:"log" yaml:"log" mapstructure:"log"`          // 日志配置
	Db     internal.Db     `json:"db" yaml:"db" mapstructure:"db"`             // 数据库配置
	Jwt    internal.Jwt    `json:"jwt" yaml:"jwt" mapstructure:"jwt"`          // Jwt配置
}
