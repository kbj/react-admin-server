package home

type ElectricityMonth struct {
	ElectricityId string  `json:"electricityId,omitempty" validate:"required" comment:"户号"`
	Month         string  `json:"month,omitempty" validate:"required,max=50" comment:"月份"`
	Amount        int     `json:"amount,omitempty" comment:"用量"`
	Fee           float64 `json:"fee,omitempty" comment:"电费"`
}

type ElectricityDay struct {
	ElectricityId string  `json:"electricityId,omitempty" validate:"required" comment:"户号"`
	Date          string  `json:"date,omitempty" comment:"日期"`
	PeakAmount    float64 `json:"peakAmount,omitempty" comment:"峰用电"`
	ValleyAmount  float64 `json:"valleyAmount,omitempty" comment:"谷用电"`
	TotalAmount   float64 `json:"totalAmount,omitempty" comment:"总用电"`
}
