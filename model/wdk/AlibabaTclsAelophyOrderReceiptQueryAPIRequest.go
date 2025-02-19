package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyOrderReceiptQueryAPIRequest 订单小票查询 API请求
// alibaba.tcls.aelophy.order.receipt.query
//
// 订单小票查询
type AlibabaTclsAelophyOrderReceiptQueryAPIRequest struct {
	model.Params
	// 小票查询请求
	_receiptQueryRequest *ReceiptQueryRequest
}

// NewAlibabaTclsAelophyOrderReceiptQueryRequest 初始化AlibabaTclsAelophyOrderReceiptQueryAPIRequest对象
func NewAlibabaTclsAelophyOrderReceiptQueryRequest() *AlibabaTclsAelophyOrderReceiptQueryAPIRequest {
	return &AlibabaTclsAelophyOrderReceiptQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyOrderReceiptQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.order.receipt.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyOrderReceiptQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReceiptQueryRequest is ReceiptQueryRequest Setter
// 小票查询请求
func (r *AlibabaTclsAelophyOrderReceiptQueryAPIRequest) SetReceiptQueryRequest(_receiptQueryRequest *ReceiptQueryRequest) error {
	r._receiptQueryRequest = _receiptQueryRequest
	r.Set("receipt_query_request", _receiptQueryRequest)
	return nil
}

// GetReceiptQueryRequest ReceiptQueryRequest Getter
func (r AlibabaTclsAelophyOrderReceiptQueryAPIRequest) GetReceiptQueryRequest() *ReceiptQueryRequest {
	return r._receiptQueryRequest
}
