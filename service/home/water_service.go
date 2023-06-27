package home

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"react-admin-server/entity/domain"
	"react-admin-server/global/g"
	"react-admin-server/tool"
	"react-admin-server/tool/r"
	"time"
)

type WaterService struct {
}

// Add 新增
func (*WaterService) Add(ctx *fiber.Ctx, param *[]domain.Water) error {
	return g.DbClient.Transaction(func(tx *gorm.DB) error {
		var count int64
		for _, water := range *param {
			// 参数校验
			if err := tool.ValidateParams(&water); err != nil {
				return err
			}
			// 检查对应的户号和时间是否重复
			if err := tx.Model(&domain.Water{}).Where("water_id = ? and cost_date = ?", water.WaterId, water.CostDate).Count(&count).Error; err != nil {
				return err
			} else if count > 0 {
				continue
			}

			// 保存数据库
			water.ID = 0
			water.CreateBy = 1
			water.CreateAt = time.Now().UnixMilli()
			if err := tx.Save(&water).Error; err != nil {
				return err
			}
		}
		return r.Ok(ctx, r.Msg("保存成功"))
	})
}
