// Package internal 系统信息配置
// @author: kbj
// @date: 2023/1/30
package internal

// System 系统信息配置
type System struct {
	Host        string `json:"host" yaml:"host" mapstructure:"host"`                        // 监听地址
	Port        string `json:"port" yaml:"port" mapstructure:"port"`                        // 监听端口
	ProjectName string `json:"projectName" yaml:"project-name" mapstructure:"project-name"` // 项目名称
	BodyLimit   int    `json:"bodyLimit" yaml:"body-limit" mapstructure:"body-limit"`       // 最大传输大小（单位MB）
	UploadPath  string `json:"uploadPath" yaml:"upload-path" mapstructure:"upload-path"`    // 附件上传路径
}
