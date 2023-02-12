// Package cmd 自动建表
// @author: kbj
// @date: 2023/2/10
package cmd

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"react-admin-server/entity/domain"
	"react-admin-server/global/g"
)

func initTables(db *gorm.DB) {
	err := db.AutoMigrate(
		domain.User{},
		domain.Role{},
		domain.Menu{},
		domain.Dept{},
	)
	if err != nil {
		g.Logger.Error("自动建表失败", zap.Error(err))
	} else {
		g.Logger.Info("自动建表成功")
	}
}
