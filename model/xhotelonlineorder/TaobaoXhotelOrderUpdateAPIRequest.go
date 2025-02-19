package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderUpdateAPIRequest 酒店订单发货接口 API请求
// taobao.xhotel.order.update
//
// 卖家确认订单或者取消订单，适用于预付、面付、信用住订单
type TaobaoXhotelOrderUpdateAPIRequest struct {
	model.Params
	// 是否把代理直签的订单同步到酒店，Y为同步，N不同步
	_syncToHotel string
	// 外部订单号（酒店确认号）
	_confirmCode string
	// 订单号
	_tid int64
	// 操作的类型：1.确认无房（取消预订，710发送短信提醒买家申请退款）2.确认预订 3.入住 4.离店 5.noshow 6.关单
	_optType int64
	// 退款费用
	_refundFee int64
	// 取消类型，6 代表的是用户取消，reasonType=7代表的是小二协商
	_reasonType int64
	// 开票金额
	_invoiceAmount int64
}

// NewTaobaoXhotelOrderUpdateRequest 初始化TaobaoXhotelOrderUpdateAPIRequest对象
func NewTaobaoXhotelOrderUpdateRequest() *TaobaoXhotelOrderUpdateAPIRequest {
	return &TaobaoXhotelOrderUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSyncToHotel is SyncToHotel Setter
// 是否把代理直签的订单同步到酒店，Y为同步，N不同步
func (r *TaobaoXhotelOrderUpdateAPIRequest) SetSyncToHotel(_syncToHotel string) error {
	r._syncToHotel = _syncToHotel
	r.Set("sync_to_hotel", _syncToHotel)
	return nil
}

// GetSyncToHotel SyncToHotel Getter
func (r TaobaoXhotelOrderUpdateAPIRequest) GetSyncToHotel() string {
	return r._syncToHotel
}

// SetConfirmCode is ConfirmCode Setter
// 外部订单号（酒店确认号）
func (r *TaobaoXhotelOrderUpdateAPIRequest) SetConfirmCode(_confirmCode string) error {
	r._confirmCode = _confirmCode
	r.Set("confirm_code", _confirmCode)
	return nil
}

// GetConfirmCode ConfirmCode Getter
func (r TaobaoXhotelOrderUpdateAPIRequest) GetConfirmCode() string {
	return r._confirmCode
}

// SetTid is Tid Setter
// 订单号
func (r *TaobaoXhotelOrderUpdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoXhotelOrderUpdateAPIRequest) GetTid() int64 {
	return r._tid
}

// SetOptType is OptType Setter
// 操作的类型：1.确认无房（取消预订，710发送短信提醒买家申请退款）2.确认预订 3.入住 4.离店 5.noshow 6.关单
func (r *TaobaoXhotelOrderUpdateAPIRequest) SetOptType(_optType int64) error {
	r._optType = _optType
	r.Set("opt_type", _optType)
	return nil
}

// GetOptType OptType Getter
func (r TaobaoXhotelOrderUpdateAPIRequest) GetOptType() int64 {
	return r._optType
}

// SetRefundFee is RefundFee Setter
// 退款费用
func (r *TaobaoXhotelOrderUpdateAPIRequest) SetRefundFee(_refundFee int64) error {
	r._refundFee = _refundFee
	r.Set("refund_fee", _refundFee)
	return nil
}

// GetRefundFee RefundFee Getter
func (r TaobaoXhotelOrderUpdateAPIRequest) GetRefundFee() int64 {
	return r._refundFee
}

// SetReasonType is ReasonType Setter
// 取消类型，6 代表的是用户取消，reasonType=7代表的是小二协商
func (r *TaobaoXhotelOrderUpdateAPIRequest) SetReasonType(_reasonType int64) error {
	r._reasonType = _reasonType
	r.Set("reason_type", _reasonType)
	return nil
}

// GetReasonType ReasonType Getter
func (r TaobaoXhotelOrderUpdateAPIRequest) GetReasonType() int64 {
	return r._reasonType
}

// SetInvoiceAmount is InvoiceAmount Setter
// 开票金额
func (r *TaobaoXhotelOrderUpdateAPIRequest) SetInvoiceAmount(_invoiceAmount int64) error {
	r._invoiceAmount = _invoiceAmount
	r.Set("invoice_amount", _invoiceAmount)
	return nil
}

// GetInvoiceAmount InvoiceAmount Getter
func (r TaobaoXhotelOrderUpdateAPIRequest) GetInvoiceAmount() int64 {
	return r._invoiceAmount
}
