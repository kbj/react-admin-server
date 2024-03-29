// Package consts 常量定义
// @author: kbj
// @date: 2023/1/30
package consts

// 配置文件相关常量
const (
	DefaultConfigEnvPath  = "RAS_CONFIG_PATH" // 在系统环境变量中配置系统配置文件路径字段的KEY
	DefaultConfigPath     = "./"              // 默认配置文件位置
	DefaultConfigFileName = "config"          // 默认配置文件名称
	DefaultConfigFileType = "yaml"            // 默认配置文件类型
)

// 状态
const (
	StatusSuccess   = 0    // 成功
	StatusFailure   = -1   // 失败
	StatusNeedLogin = 4000 // 需要登录
)

// 分页
const (
	PagingPageNum  = "1"
	PagingPageSize = "10"
)

// 文件
const (
	MaxFileNameSize = 200 // 最大文件名长度大小
)

// FileAllowedExtension 默认允许文件格式
var FileAllowedExtension = [...]string{
	// 图片
	".bmp", ".gif", ".jpg", ".jpeg", ".png",
	// word excel powerpoint
	".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt",
	// 压缩文件
	".rar", ".zip", ".gz", ".bz2",
	// 视频格式
	".mp4", ".avi", ".rmvb",
	// pdf
	".pdf",
}
