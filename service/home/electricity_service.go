package home

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"react-admin-server/entity/domain"
	"react-admin-server/entity/vo/home"
	"react-admin-server/global/g"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
	"time"
)

type ElectricityService struct {
}

// AddMonth 月电量统计新增
func (*ElectricityService) AddMonth(ctx *fiber.Ctx, param *[]home.ElectricityMonth) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		var count int64
		for _, month := range *param {
			// 参数校验
			if err := tool.ValidateParams(&month); err != nil {
				return err
			}

			// 检查月份是否存在
			if err := tx.Model(&domain.ElectricityMonth{}).Where("month = ?", month.Month).Count(&count).Error; err != nil {
				return err
			} else if count > 0 {
				continue
			}

			// 保存数据库
			entity := domain.ElectricityMonth{
				Common: domain.Common{
					CreateAt: time.Now().UnixMilli(),
					CreateBy: 1,
				},
				Month:  month.Month,
				Amount: month.Amount,
				Fee:    month.Fee,
			}
			if err := tx.Save(&entity).Error; err != nil {
				return err
			}
		}

		return r.Ok(ctx, r.Msg("保存成功"))
	})
}

// AddDay 日电量统计新增
func (*ElectricityService) AddDay(ctx *fiber.Ctx, param *[]home.ElectricityDay) error {
	location, _ := time.LoadLocation("Asia/Shanghai")

	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		var count int64
		for _, day := range *param {
			// 解析时间
			date, err := time.ParseInLocation("2006-01-02", day.Date, location)
			if err != nil {
				return err
			}

			// 检查日期是否存在
			if err = tx.Model(&domain.ElectricityDay{}).Where("date = ?", date).Count(&count).Error; err != nil {
				return err
			} else if count > 0 {
				continue
			}

			// 保存数据库
			entity := domain.ElectricityDay{
				Common: domain.Common{
					CreateAt: time.Now().UnixMilli(),
					CreateBy: 1,
				},
				Date:         date,
				PeakAmount:   day.PeakAmount,
				ValleyAmount: day.ValleyAmount,
				TotalAmount:  day.TotalAmount,
			}
			if err = tx.Save(&entity).Error; err != nil {
				return err
			}
		}
		return r.Ok(ctx, r.Msg("保存成功"))
	})
}
