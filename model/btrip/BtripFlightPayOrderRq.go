package btrip

// BtripFlightPayOrderRq 结构体
type BtripFlightPayOrderRq struct {
	// 分销外部订单号
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 扩展信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 分销子渠道，通常为corpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 订单支付总金额
	TotalPayPrice int64 `json:"total_pay_price,omitempty" xml:"total_pay_price,omitempty"`
	// 企业支付金额
	CorpPayPrice int64 `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	// 个人支付金额
	PersonalPayPrice int64 `json:"personal_pay_price,omitempty" xml:"personal_pay_price,omitempty"`
	// 商旅订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
