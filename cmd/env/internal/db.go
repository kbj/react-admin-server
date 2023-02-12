// Package internal 数据库配置
// @author: kbj
// @date: 2023/1/31
package internal

type Db struct {
	DbType      string `json:"dbType" yaml:"db-type" mapstructure:"dbType"`                   // 数据库类型
	Host        string `json:"host" yaml:"host" mapstructure:"host"`                          // 地址
	Port        string `json:"port" yaml:"port" mapstructure:"port"`                          // 端口
	User        string `json:"user" yaml:"user" mapstructure:"user"`                          // 用户名
	Password    string `json:"password" yaml:"password" mapstructure:"password"`              // 密码
	DbName      string `json:"dbName" yaml:"db-name" mapstructure:"dbName"`                   // 数据库名称
	Params      string `json:"params" yaml:"params" mapstructure:"params"`                    // 连接参数
	MaxIdleConn int    `json:"maxIdleConn" yaml:"max-idle-conn" mapstructure:"max-idle-conn"` // 空闲中的最大连接数
	MaxOpenConn int    `json:"maxOpenConn" yaml:"max-open-conn" mapstructure:"max-open-conn"` // 打开到数据库的最大连接数
	PrintSQL    bool   `json:"printSQL" yaml:"print-sql" mapstructure:"print-sql"`            // 打印执行SQL
}
