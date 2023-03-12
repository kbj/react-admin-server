// Package dao 用户表相关DAO
// @author: kbj
// @date: 2023/3/12
package dao

import (
	"gorm.io/gorm"
	"react-admin-server/entity/domain"
	"react-admin-server/entity/vo/system"
	"react-admin-server/tool"
)

type UserDao struct {
}

func (*UserDao) SelectUserList(db *gorm.DB, result *[]domain.User, param *system.UserSearch) error {
	db = db.Model(&domain.User{})

	// 查询条件
	if param.Username != "" {
		db.Where("username like ?", "%"+param.Username+"%")
	}
	if param.NickName != "" {
		db.Where("nick_name like ?", "%"+param.NickName+"%")
	}
	if param.Mobile != "" {
		db.Where("mobile = ?", param.Mobile)
	}
	if param.Gender != "" {
		db.Where("gender = ?", param.Gender)
	}
	return tool.LogDbError(db.Find(result).Error)
}

// SelectLoginUser 根据登录凭证查询
func (*UserDao) SelectLoginUser(db *gorm.DB, result *domain.User, username string, password string) error {
	return db.Where("username = ? and password = ?", username, password).First(result).Error
}
