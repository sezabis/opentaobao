package flight

// TicketingDetailDto 结构体
type TicketingDetailDto struct {
	// 出票对象
	IssueList []IssueList `json:"issue_list,omitempty" xml:"issue_list>issue_list,omitempty"`
	// 订单标签
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 飞猪订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 出票时间
	IssueTime string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// sla
	Sla string `json:"sla,omitempty" xml:"sla,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 退改签规则
	RefundModifyRule string `json:"refund_modify_rule,omitempty" xml:"refund_modify_rule,omitempty"`
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 国内国际标识
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
	// 佣金
	Commission int64 `json:"commission,omitempty" xml:"commission,omitempty"`
	// 订单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
