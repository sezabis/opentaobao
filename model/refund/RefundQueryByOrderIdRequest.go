package refund

// RefundQueryByOrderIdRequest 结构体
type RefundQueryByOrderIdRequest struct {
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
