package aesolution

// GlobalAeopTpOrderDetailDto 结构体
type GlobalAeopTpOrderDetailDto struct {
	// Order Message list(deprecated)
	OrderMsgList []GlobalAeopTpOrderMsgDto `json:"order_msg_list,omitempty" xml:"order_msg_list>global_aeop_tp_order_msg_dto,omitempty"`
	// child order ext info list
	ChildOrderExtInfoList []GlobalAeopTpOrderProductInfoDto `json:"child_order_ext_info_list,omitempty" xml:"child_order_ext_info_list>global_aeop_tp_order_product_info_dto,omitempty"`
	// logistics info
	LogisticInfoList []GlobalAeopTpLogisticInfoDto `json:"logistic_info_list,omitempty" xml:"logistic_info_list>global_aeop_tp_logistic_info_dto,omitempty"`
	// operation details list
	OprLogDtoList []GlobalAeopTpOperationLogDto `json:"opr_log_dto_list,omitempty" xml:"opr_log_dto_list>global_aeop_tp_operation_log_dto,omitempty"`
	// child order list
	ChildOrderList []GlobalAeopTpChildOrderDto `json:"child_order_list,omitempty" xml:"child_order_list>global_aeop_tp_child_order_dto,omitempty"`
	// modified time, it&#39;s US pacific time
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// Order end time
	GmtTradeEnd string `json:"gmt_trade_end,omitempty" xml:"gmt_trade_end,omitempty"`
	// buyer login id
	Buyerloginid string `json:"buyerloginid,omitempty" xml:"buyerloginid,omitempty"`
	// buyer memo
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// logisitcs escrow fee rate(Deprecated)
	LogisitcsEscrowFeeRate string `json:"logisitcs_escrow_fee_rate,omitempty" xml:"logisitcs_escrow_fee_rate,omitempty"`
	// Payment settlement currency
	SettlementCurrency string `json:"settlement_currency,omitempty" xml:"settlement_currency,omitempty"`
	// order pay amount(settlemet currency)
	PayAmountBySettlementCur string `json:"pay_amount_by_settlement_cur,omitempty" xml:"pay_amount_by_settlement_cur,omitempty"`
	// frozen status
	FrozenStatus string `json:"frozen_status,omitempty" xml:"frozen_status,omitempty"`
	// issue status
	IssueStatus string `json:"issue_status,omitempty" xml:"issue_status,omitempty"`
	// logistics status：NO_LOGISTICS 、 WAIT_SELLER_SEND_GOODS, SELLER_SEND_PART_GOODS, SELLER_SEND_GOODS, BUYER_ACCEPT_GOODS,NO_LOGISTICS
	LogisticsStatus string `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
	// Seller full name
	SellerSignerFullname string `json:"seller_signer_fullname,omitempty" xml:"seller_signer_fullname,omitempty"`
	// Current status expiration date
	OverTimeLeft string `json:"over_time_left,omitempty" xml:"over_time_left,omitempty"`
	// order end reason
	OrderEndReason string `json:"order_end_reason,omitempty" xml:"order_end_reason,omitempty"`
	// order creation time
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// seller operator login ID
	SellerOperatorLoginId string `json:"seller_operator_login_id,omitempty" xml:"seller_operator_login_id,omitempty"`
	// payment type
	PaymentType string `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// Order Status：PLACE_ORDER_SUCCESS;  IN_CANCEL;  WAIT_SELLER_SEND_GOODS;  SELLER_PART_SEND_GOODS;  WAIT_BUYER_ACCEPT_GOODS;  FUND_PROCESSING; IN_ISSUE;  IN_FROZEN;  WAIT_SELLER_EXAMINE_MONEY;  RISK_CONTROL.
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// buyer full name
	BuyerSignerFullname string `json:"buyer_signer_fullname,omitempty" xml:"buyer_signer_fullname,omitempty"`
	// successful payment time
	GmtPaySuccess string `json:"gmt_pay_success,omitempty" xml:"gmt_pay_success,omitempty"`
	// loan status
	LoanStatus string `json:"loan_status,omitempty" xml:"loan_status,omitempty"`
	// seller operator ali login id
	SellerOperatorAliidloginid string `json:"seller_operator_aliidloginid,omitempty" xml:"seller_operator_aliidloginid,omitempty"`
	// fund status
	FundStatus string `json:"fund_status,omitempty" xml:"fund_status,omitempty"`
	// cpf  number of order
	CpfNumber string `json:"cpf_number,omitempty" xml:"cpf_number,omitempty"`
	// pick-up type of order
	OfflinePickupType string `json:"offline_pickup_type,omitempty" xml:"offline_pickup_type,omitempty"`
	// pickup point code of the specific order
	OfflinePickupPointCode string `json:"offline_pickup_point_code,omitempty" xml:"offline_pickup_point_code,omitempty"`
	// buyer info
	BuyerInfo *GlobalAeopTpPersonDto `json:"buyer_info,omitempty" xml:"buyer_info,omitempty"`
	// receipt address
	ReceiptAddress *GlobalAeopTpAddressDto `json:"receipt_address,omitempty" xml:"receipt_address,omitempty"`
	// logistics amount
	LogisticsAmount *GlobalMoneyStr `json:"logistics_amount,omitempty" xml:"logistics_amount,omitempty"`
	// issue info
	IssueInfo *GlobalAeopTpIssueInfoDto `json:"issue_info,omitempty" xml:"issue_info,omitempty"`
	// refund info
	RefundInfo *GlobalAeopTpRefundInfoDto `json:"refund_info,omitempty" xml:"refund_info,omitempty"`
	// order ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// order amount
	OrderAmount *GlobalMoneyStr `json:"order_amount,omitempty" xml:"order_amount,omitempty"`
	// order amount
	InitOderAmount *GlobalMoneyStr `json:"init_oder_amount,omitempty" xml:"init_oder_amount,omitempty"`
	// loan info
	LoanInfo *GlobalAeopTpLoanInfoDto `json:"loan_info,omitempty" xml:"loan_info,omitempty"`
	// escrow fee (deprecated)
	EscrowFee *GlobalMoneyStr `json:"escrow_fee,omitempty" xml:"escrow_fee,omitempty"`
	// Order discount total amount (sum of the platform and seller discounts)
	OrderDiscountInfo *GlobalMoneyStr `json:"order_discount_info,omitempty" xml:"order_discount_info,omitempty"`
	// phone order or not
	IsPhone bool `json:"is_phone,omitempty" xml:"is_phone,omitempty"`
}
