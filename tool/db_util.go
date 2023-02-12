// Package tool 数据库error日志打印
// @author: kbj
// @date: 2023/2/10
package tool

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"react-admin-server/entity/vo"
	"react-admin-server/global/consts"
	"react-admin-server/global/enums"
	"react-admin-server/global/g"
	"strconv"
)

func LogDbError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) || err == nil {
		return nil
	}
	g.Logger.Error("数据库执行失败", zap.Error(err))
	return err
}

// SelectPageList 分页查询
func SelectPageList[T any](ctx *fiber.Ctx, db *gorm.DB) (*vo.Page[*T], error) {
	var page vo.Page[*T]
	// 页码
	pageNum, _ := strconv.Atoi(ctx.Query(enums.PagingPageNum, consts.PagingPageNum))
	if pageNum == 0 {
		pageNum = 1
	}
	page.PageNum = pageNum

	// 每页大小
	pageSize, _ := strconv.Atoi(ctx.Query(enums.PagingPageSize, consts.PagingPageSize))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	page.PageSize = pageSize

	// 排序规则 todo 还需要防止SQL注入
	orderBy := ctx.Query(enums.PagingOrderBy)
	if orderBy != "" {
		isDesc := ctx.Query(enums.PagingIsDesc, "false")
		if isDesc == "true" {
			orderBy = orderBy + " desc"
		}
	}

	// 记录数
	db.Count(&page.Total)
	if page.Total > 0 {
		return &page, LogDbError(db.Scopes(paginate(pageNum, pageSize, orderBy)).Find(&page.Records).Error)
	}
	return &page, nil
}

// paginate Gorm分页中间件
func paginate(pageNum int, pageSize int, orderBy string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageNum - 1) * pageSize
		if orderBy != "" {
			return db.Order(orderBy).Offset(offset).Limit(pageSize)
		} else {
			return db.Offset(offset).Limit(pageSize)
		}
	}
}
