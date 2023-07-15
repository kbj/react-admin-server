package domain

import "time"

type ElectricityMonth struct {
	Common
	ElectricityId string  `json:"electricityId,omitempty" gorm:"comment:户号;size:200;not null"`
	Month         string  `json:"month,omitempty" gorm:"comment:月份;size:50;not null"`
	Amount        int     `json:"amount,omitempty" gorm:"comment:用量"`
	Fee           float64 `json:"fee,omitempty" gorm:"comment:电费"`
}

type ElectricityDay struct {
	Common
	ElectricityId string    `json:"electricityId,omitempty" gorm:"comment:户号;size:200;not null"`
	Date          time.Time `json:"date,omitempty" gorm:"comment:日期;not null"`
	PeakAmount    float64   `json:"peakAmount,omitempty" gorm:"comment:峰用电"`
	ValleyAmount  float64   `json:"valleyAmount,omitempty" gorm:"comment:谷用电"`
	TotalAmount   float64   `json:"totalAmount,omitempty" gorm:"comment:总用电"`
}
