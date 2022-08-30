package wdk

// RefundCsApplyRenderDto 结构体
type RefundCsApplyRenderDto struct {
	// 申请退款的子订单ID列表
	OutSubOrderIds []string `json:"out_sub_order_ids,omitempty" xml:"out_sub_order_ids>string,omitempty"`
	// 渠道订单ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 商家经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道来源
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}
