// Package internal 日志配置
// @author: kbj
// @date: 2023/1/30
package internal

type Log struct {
	Path         string `json:"path" yaml:"path" mapstructure:"path"`                             // 写入位置
	ConsoleLevel string `json:"consoleLevel" yaml:"console-level" mapstructure:"console-level"`   // 控制台打印级别
	FileLevel    string `json:"fileLevel" yaml:"file-level" mapstructure:"file-level"`            // 文件写入级别
	LogInConsole bool   `json:"logInConsole" yaml:"log-in-console" mapstructure:"log-in-console"` // 控制台打印日志
	Compress     bool   `json:"compress" yaml:"compress" mapstructure:"compress"`                 // 压缩日志
}
