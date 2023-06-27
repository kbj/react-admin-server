package domain

import "react-admin-server/global/types"

type Water struct {
	Common
	WaterId          string         `json:"waterId,omitempty" gorm:"comment:水费户号;not null;size:100" validate:"required" comment:"水费户号"`
	Address          string         `json:"address,omitempty" gorm:"comment:地址;size=200"`
	Caliber          string         `json:"caliber,omitempty" gorm:"comment:口径（估计是水管直径之类的玩意）;size=100"`
	Consumption      float64        `json:"consumption,omitempty" gorm:"comment:用水量"`
	CostDate         string         `json:"costDate,omitempty" gorm:"comment:费用日期" validate:"required" comment:"费用日期"`
	CurrentRead      float64        `json:"currentRead,omitempty" gorm:"comment:当前读数" comment:"当前读数"`
	BeforeRead       float64        `json:"beforeRead,omitempty" gorm:"comment:之前读数" comment:"之前读数"`
	InputTime        types.DateTime `json:"inputTime,omitempty" gorm:"comment:抄表时间;type:time"`
	WaterFee         float64        `json:"waterFee,omitempty" gorm:"comment:水费"`
	PenaltyFee       float64        `json:"penaltyFee,omitempty" gorm:"comment:违约金"`
	GarbageFee       float64        `json:"garbageFee,omitempty" gorm:"comment:垃圾清理费"`
	TotalFee         float64        `json:"totalFee,omitempty" gorm:"comment:总账单"`
	PriceDescription string         `json:"priceDescription,omitempty" gorm:"comment:计费规则描述;size:500"`
}
