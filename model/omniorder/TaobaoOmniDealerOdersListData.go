package omniorder

// TaobaoOmniDealerOdersListData 结构体
type TaobaoOmniDealerOdersListData struct {
	// 外部系统订单id
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 主订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
