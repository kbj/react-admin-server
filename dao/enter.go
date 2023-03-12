// Package dao DAO统一入口
// @author: kbj
// @date: 2023/3/12
package dao

import "react-admin-server/dao/internal/dao"

var (
	User = new(dao.UserDao)
)
