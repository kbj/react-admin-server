// Package cmd Zap日志库
// @author: kbj
// @date: 2023/1/30
package cmd

import (
	"github.com/arthurkiller/rollingwriter"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"react-admin-server/global/g"
	"time"
)

func initZap() *zap.Logger {
	// 先检查路径是否存在，不存在创建对应文件夹
	if !fsutil.DirExist(g.Env.Log.Path) {
		log.Println("日志文件夹：" + g.Env.Log.Path + " 不存在，自动创建中...")
		if err := fsutil.MkParentDir(g.Env.Log.Path); err != nil {
			log.Fatalf("创建日志文件夹失败！%v\n", err)
		}
	}

	/*
		使用zap.New(…)自定义所有配置信息。func New(core zapcore.Core, options ...Option) *Logger

		zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel
			- Encoder: 编码器 (写入日志格式)
			- WriterSyncer ：指定日志写到哪里去
			- Log Level：哪种级别的日志将被写入
	*/
	var cores []zapcore.Core
	infoFileCore := zapcore.NewCore(getFileEncoder(), getFileWriter(g.Env.System.ProjectName+"-info"), parseLogLevel(g.Env.Log.FileLevel, true)) // error以下级别打印的日志文件
	errorFileCore := zapcore.NewCore(getFileEncoder(), getFileWriter(g.Env.System.ProjectName+"-error"), parseLogLevel("error", false))          // error及以上级别打印的日志文件
	cores = append(cores, infoFileCore, errorFileCore)
	if g.Env.Log.LogInConsole {
		// 启用控制台日志
		consoleCore := zapcore.NewCore(getConsoleEncoder(), getConsoleWriter(), parseLogLevel(g.Env.Log.ConsoleLevel, false)) // 控制台打印日志
		cores = append(cores, consoleCore)
	}
	return zap.New(zapcore.NewTee(cores...), zap.AddCaller()) // AddCaller()表示显示文件路径和行号
}

// 创建控制台的Encoder
func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 日志时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 日志级别按大写展示并且不同级别有不同颜色
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 创建控制台写入的Writer
func getConsoleWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

// 创建文件的Encoder
func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 日志时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 日志级别按大写展示
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 创建文件
func getFileWriter(fileName string) zapcore.WriteSyncer {
	rollConfig := rollingwriter.NewDefaultConfig()
	rollConfig.LogPath = g.Env.Log.Path      // 日志文件夹
	rollConfig.FileName = fileName           // 日志文件名
	rollConfig.MaxRemain = 365               // 日志保留365个文件
	rollConfig.Compress = g.Env.Log.Compress // 是否开启日志压缩

	// 创建 writer
	w, err := rollingwriter.NewWriterFromConfig(&rollConfig)
	if err != nil {
		log.Fatalf("zap writer %s 创建失败：%v\n", fileName, err)
	}

	return zapcore.AddSync(w)
}

// 将日志配置的级别解析为zap的日志级别
func parseLogLevel(level string, belowError bool) zap.LevelEnablerFunc {
	var minLevel zapcore.Level
	if err := minLevel.UnmarshalText(strutil.ToBytes(level)); err != nil {
		log.Printf("%s 日志级别解析失败，使用默认Debug级别", level)
		minLevel = zapcore.DebugLevel
	}

	finalLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		if belowError {
			return lev < zapcore.ErrorLevel && lev >= minLevel
		} else {
			return lev >= minLevel
		}
	})
	return finalLevel
}

// customEncodeTime 自定义日志输出时间格式
func customEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 / 15:04:05.000 / -0700"))
}
