package tmallgenie

// PurchaseReceiveReturnDto 结构体
type PurchaseReceiveReturnDto struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 其他数据
	Map string `json:"map,omitempty" xml:"map,omitempty"`
	// 成功状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
