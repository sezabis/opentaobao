package fenxiao

// ProductCat 结构体
type ProductCat struct {
	// 产品线名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 产品最低零售价
	RetailLowPercent string `json:"retail_low_percent,omitempty" xml:"retail_low_percent,omitempty"`
	// 产品最高零售价
	RetailHighPercent string `json:"retail_high_percent,omitempty" xml:"retail_high_percent,omitempty"`
	// 产品代销采购价
	CostPercentAgent string `json:"cost_percent_agent,omitempty" xml:"cost_percent_agent,omitempty"`
	// 产品经销采购价
	CostPercentDealer string `json:"cost_percent_dealer,omitempty" xml:"cost_percent_dealer,omitempty"`
	// 产品线ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 产品数量
	ProductNum int64 `json:"product_num,omitempty" xml:"product_num,omitempty"`
}
