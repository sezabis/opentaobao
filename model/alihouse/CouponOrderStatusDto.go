package alihouse

// CouponOrderStatusDto 结构体
type CouponOrderStatusDto struct {
	// 淘宝订单id
	TbOrderId string `json:"tb_order_id,omitempty" xml:"tb_order_id,omitempty"`
	// 完结时间
	FinishTime string `json:"finish_time,omitempty" xml:"finish_time,omitempty"`
	// code码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 履约单订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 合同单id
	ContractOrderId int64 `json:"contract_order_id,omitempty" xml:"contract_order_id,omitempty"`
	// 电商履约单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 电商履约单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 取消状态
	CancelStatus int64 `json:"cancel_status,omitempty" xml:"cancel_status,omitempty"`
}
