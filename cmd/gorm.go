// Package cmd Gorm驱动
// @author: kbj
// @date: 2023/2/10
package cmd

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"react-admin-server/global/g"
	"strings"
	"time"
)

func initGorm() *gorm.DB {
	// 数据库连接配置信息
	var dbConfig *gorm.Dialector
	switch strings.ToLower(g.Env.Db.DbType) {
	case "mysql":
		dbConfig = initGormMySQLConfig()
	case "postgres":
		dbConfig = initGormPGSQLConfig()
	default:
		dbConfig = initGormMySQLConfig()
	}

	// Gorm的配置
	gormConfig := gorm.Config{
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀
			SingularTable: true, // 表名用单数名词
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
	}

	// 启用日志
	logger := zapgorm2.New(g.Logger)
	logger.SetAsDefault()
	logger.IgnoreRecordNotFoundError = true
	if g.Env.Db.PrintSQL {
		// info级别即可打印所有日志
		logger.LogLevel = gormLogger.Info
	}
	gormConfig.Logger = logger

	// 建立连接
	db, err := gorm.Open(*dbConfig, &gormConfig)
	if err != nil {
		g.Logger.Fatal("数据库连接建立失败", zap.Error(err))
	}

	// 设置连接池信息
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(g.Env.Db.MaxOpenConn)
	sqlDB.SetMaxIdleConns(g.Env.Db.MaxIdleConn)

	// 自动建表
	initTables(db)
	return db
}

// 初始化MySQL数据库连接配置
func initGormMySQLConfig() *gorm.Dialector {
	mysqlConfig := mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", g.Env.Db.User, g.Env.Db.Password, g.Env.Db.Host, g.Env.Db.Port, g.Env.Db.DbName, g.Env.Db.Params),
		SkipInitializeWithVersion: false,
		DefaultStringSize:         200,
	}
	config := mysql.New(mysqlConfig)
	return &config
}

// 初始化PG数据库连接配置
func initGormPGSQLConfig() *gorm.Dialector {
	pgConfig := postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", g.Env.Db.Host, g.Env.Db.User, g.Env.Db.Password, g.Env.Db.DbName, g.Env.Db.Port),
		PreferSimpleProtocol: false,
	}
	config := postgres.New(pgConfig)
	return &config
}
