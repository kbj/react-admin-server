// Package cmd 初始化数据库连接
// @author: kbj
// @date: 2023/1/31
package cmd

//
//import (
//	"context"
//	"entgo.io/ent/dialect"
//	"entgo.io/ent/dialect/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/spf13/cast"
//	"go.uber.org/zap"
//	"react-admin-server/ent"
//	"react-admin-server/ent/migrate"
//	"react-admin-server/global/g"
//	"time"
//)
//
//func initEnt() (client *ent.Client) {
//	driver, err := sql.Open(g.Env.Db.DbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", g.Env.Db.User, g.Env.Db.Password, g.Env.Db.Host, g.Env.Db.Port, g.Env.Db.DbName, g.Env.Db.Params))
//	if err != nil {
//		g.Logger.Fatal("数据库连接初始化失败", zap.Error(err))
//	}
//	db := driver.DB()
//	db.SetConnMaxLifetime(time.Hour)
//	db.SetMaxOpenConns(g.Env.Db.MaxOpenConn)
//	db.SetMaxIdleConns(g.Env.Db.MaxIdleConn)
//
//	// 启用日志打印
//	if g.Env.Db.PrintSQL {
//		driverWithContext := dialect.DebugWithContext(driver, func(ctx context.Context, a ...any) {
//			// 打印执行SQL
//			if len(a) == 1 {
//				g.Logger.Info(cast.ToString(a[0]))
//			} else {
//				g.Logger.Info("", zap.Any("exec sql:", a))
//			}
//		})
//		client = ent.NewClient(ent.Driver(driverWithContext))
//	} else {
//		client = ent.NewClient(ent.Driver(driver))
//	}
//
//	// 自动同步表结构
//	if err = client.Schema.Create(
//		context.Background(),
//		migrate.WithForeignKeys(false), // 不带外键
//		migrate.WithDropIndex(true),    // 删除不存在的索引
//		migrate.WithDropColumn(true),   // 删除不存在的列
//	); err != nil {
//		g.Logger.Fatal("自动同步表结构失败", zap.Error(err))
//	}
//
//	return client
//}
