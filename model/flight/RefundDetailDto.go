package flight

// RefundDetailDto 结构体
type RefundDetailDto struct {
	// 退票数据集
	RefundList []RefundList `json:"refund_list,omitempty" xml:"refund_list>refund_list,omitempty"`
	// 紧急标识:EPIDEMIC:疫情审核通过标, REASSURING_TICKET:安心票, FLIGHTCHANGE:航变审核通过标, CALCULATED:预计算, FLIGHTCHANGECANCEL:航班取消, FLIGHTCHANGEDELAY:航班延误, ILLWITHCERTIFICATE:病退, REFUSETOTAKE:航司拒载, PREVIOUSLATTERINTERFERENCE:前后段影响, OVERBOOKING:超售, AIRLINEAUTHORIZE:航司授权, CHANGENAME:改名, REPEATBUYTICKETS:重复购票, BUYWRONGTICKETS:错购, AIRLINEVIPCARDREFUND:航司金银卡退票
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 申请单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 申请原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 关联飞猪订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// sla
	Sla string `json:"sla,omitempty" xml:"sla,omitempty"`
	// 交易币种:CNY:人民币, USD:美元, HKD:港元
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 申请时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 拒绝原因
	RefuseReason string `json:"refuse_reason,omitempty" xml:"refuse_reason,omitempty"`
	// 自愿标识:0.非自愿，1.自愿
	Voluntary int64 `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 国际国内标识(1国内,2国际)
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
	// 退票状态:待回填:1, 待退款:2,	 退款中:3,	 已完结:4,	 已拒绝:5
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 申请原因类型:0:其他, 1:航变,2:病退 ,3:疫情,4:特殊提退
	ApplyReasonType int64 `json:"apply_reason_type,omitempty" xml:"apply_reason_type,omitempty"`
	// 补退单标识(1是补退单)
	Supplement int64 `json:"supplement,omitempty" xml:"supplement,omitempty"`
}
