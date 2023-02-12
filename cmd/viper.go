package cmd

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"react-admin-server/global/consts"
	"react-admin-server/global/g"
)

// initViper 初始化一个新的Viper实例
// @path 参数文件路径
func initViper(path ...string) *viper.Viper {
	// 用命令行指定配置文件的路径
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "配置文件路径")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(consts.DefaultConfigEnvPath); configEnv == "" {
				config = consts.DefaultConfigPath + consts.DefaultConfigFileName + "." + consts.DefaultConfigFileType
				log.Printf("您正在使用配置文件路径的默认值，config的路径为%v\n", config)
			} else {
				config = configEnv
				log.Printf("您正在使用%s环境变量，配置文件的路径为%v\n", configEnv, config)
			}
		} else {
			log.Printf("您正在使用命令行的-c参数传递的值，配置文件的路径为%v\n", config)
		}
	} else {
		config = path[0]
		log.Printf("您正在使用初始化函数传递的值，配置文件的路径为%v\n", config)
	}

	// 新建viper实例
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType(consts.DefaultConfigFileType)
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %s \n", err)
	}

	// 添加默认值
	addDefault(v)

	// 将读取到的配置信息反序列化到配置文件单例
	if err = v.Unmarshal(&g.Env); err != nil {
		log.Println(err)
	}

	// 配置文件热更新
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件已变更:" + e.Name)
		// 更新配置文件单例
		if err = v.Unmarshal(&g.Env); err != nil {
			log.Printf("配置文件热更新失败：%v\n", err)
		}
	})

	return v
}

// 添加默认设置
func addDefault(v *viper.Viper) {
	v.SetDefault("system.host", "0.0.0.0")                    // 监听地址
	v.SetDefault("system.port", "8080")                       // 启动端口
	v.SetDefault("system.project-name", "react-admin-server") // 项目名称
	v.SetDefault("system.body-limit", 100)                    // 最大传输大小

	v.SetDefault("log.path", "./logs")         // 日志保存路径
	v.SetDefault("log.console-level", "debug") // 控制台日志级别
	v.SetDefault("log.file-level", "debug")    // 文件日志级别
	v.SetDefault("log.log-in-console", true)   // 控制台打印日志
	v.SetDefault("log.compress", true)         // 压缩日志

	v.SetDefault("db.max-open-conn", 500) // 打开到数据库的最大连接数
	v.SetDefault("db.max-idle-conn", 50)  // 空闲中的最大连接数

	v.SetDefault("jwt.expire", 60)                 // JWT过期时间分钟
	v.SetDefault("jwt.name", "react-admin-server") // JWT签发者
	v.SetDefault("jwt.key", "1234567890")          // JWT默认密钥
}
