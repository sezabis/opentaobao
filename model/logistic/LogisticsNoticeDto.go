package logistic

// LogisticsNoticeDto 结构体
type LogisticsNoticeDto struct {
	// 商品信息
	CommodityInfos []CommodityInfo `json:"commodity_infos,omitempty" xml:"commodity_infos>commodity_info,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 快递公司编码
	TpCode string `json:"tp_code,omitempty" xml:"tp_code,omitempty"`
	// 导入类型(1-订单拆包；2-赠品；3-补发；99-以上类型都不是可填写99)
	ImportType int64 `json:"import_type,omitempty" xml:"import_type,omitempty"`
	// 主交易单号
	ParentOrderId int64 `json:"parent_order_id,omitempty" xml:"parent_order_id,omitempty"`
}
