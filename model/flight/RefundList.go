package flight

// RefundList 结构体
type RefundList struct {
	// 票号,必填
	Tickets []string `json:"tickets,omitempty" xml:"tickets>string,omitempty"`
	// 退票航段
	RefundSegments []RefundSegments `json:"refund_segments,omitempty" xml:"refund_segments>refund_segments,omitempty"`
	// 税费，单位:分
	Taxes []Tax `json:"taxes,omitempty" xml:"taxes>tax,omitempty"`
	// 乘客信息,必填
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 退票费,必填,单位:分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 升舱手续费,改后退订单必填,单位:分
	RefundUpgradeFee int64 `json:"refund_upgrade_fee,omitempty" xml:"refund_upgrade_fee,omitempty"`
	// 改期手续费,改后退订单必填,单位:分
	RefundModifyFee int64 `json:"refund_modify_fee,omitempty" xml:"refund_modify_fee,omitempty"`
	// 乘客类型:1:成人, 2:儿童, 3:婴儿, 4:留学生
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 退款金额，单位:分（出后退:退款金额=实收金额-退票费；改后退:退款金额=实收金额+总改期费+总升舱费-退票费-改签手续费-升舱手续费）
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 票面价，单位:分
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 升舱总费用，单位:分
	TotalModifyFee int64 `json:"total_modify_fee,omitempty" xml:"total_modify_fee,omitempty"`
	// 改签总费用，单位:分
	TotalUpgradeFee int64 `json:"total_upgrade_fee,omitempty" xml:"total_upgrade_fee,omitempty"`
	// 实收金额：单位:分
	RealPrice int64 `json:"real_price,omitempty" xml:"real_price,omitempty"`
	// 1:出后退,2:改后退
	RefundItemType int64 `json:"refund_item_type,omitempty" xml:"refund_item_type,omitempty"`
}
