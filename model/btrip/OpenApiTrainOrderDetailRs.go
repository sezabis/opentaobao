package btrip

// OpenApiTrainOrderDetailRs 结构体
type OpenApiTrainOrderDetailRs struct {
	// 订单费用列表
	PriceInfoList []PriceInfo `json:"price_info_list,omitempty" xml:"price_info_list>price_info,omitempty"`
	// 改签票列表
	ChangeTicketInfoList []TrainChangeTicketInfo `json:"change_ticket_info_list,omitempty" xml:"change_ticket_info_list>train_change_ticket_info,omitempty"`
	// 退票列表
	RefundTicketInfoList []TrainRefundTicketInfo `json:"refund_ticket_info_list,omitempty" xml:"refund_ticket_info_list>train_refund_ticket_info,omitempty"`
	// 正票列表
	TicketInfoList []TrainTicketInfo `json:"ticket_info_list,omitempty" xml:"ticket_info_list>train_ticket_info,omitempty"`
	// 出行人列表
	PassengerInfoList []PassengerInfo `json:"passenger_info_list,omitempty" xml:"passenger_info_list>passenger_info,omitempty"`
	// 订单基础信息
	OrderBaseInfo *OrderBaseInfo `json:"order_base_info,omitempty" xml:"order_base_info,omitempty"`
	// 发票信息
	InvoiceInfo *InvoiceInfo `json:"invoice_info,omitempty" xml:"invoice_info,omitempty"`
	// 车次信息
	TrainInfo *TrainInfo `json:"train_info,omitempty" xml:"train_info,omitempty"`
}
