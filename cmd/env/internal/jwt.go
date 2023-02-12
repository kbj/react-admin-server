// Package internal JWT相关配置
// @author: kbj
// @date: 2023/2/5
package internal

type Jwt struct {
	Expire int      `json:"expire" yaml:"expire" mapstructure:"expire"` // 超时时间(单位分钟)
	Name   string   `json:"name" yaml:"name" mapstructure:"name"`       // 签发者
	Key    string   `json:"key" yaml:"key" mapstructure:"key"`          // 密钥
	ByPass []string `json:"byPass" yaml:"byPass" mapstructure:"byPass"` // 绕过校验路径
}
