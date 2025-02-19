package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBcFutureStockInboundAPIRequest 供应商上报期货库存 API请求
// alibaba.alihealth.bc.future.stock.inbound
//
// 供应商上报期货库存
type AlibabaAlihealthBcFutureStockInboundAPIRequest struct {
	model.Params
	// 请求体
	_futureInboundReqDto *FutureInboundReqDto
}

// NewAlibabaAlihealthBcFutureStockInboundRequest 初始化AlibabaAlihealthBcFutureStockInboundAPIRequest对象
func NewAlibabaAlihealthBcFutureStockInboundRequest() *AlibabaAlihealthBcFutureStockInboundAPIRequest {
	return &AlibabaAlihealthBcFutureStockInboundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBcFutureStockInboundAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.bc.future.stock.inbound"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBcFutureStockInboundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFutureInboundReqDto is FutureInboundReqDto Setter
// 请求体
func (r *AlibabaAlihealthBcFutureStockInboundAPIRequest) SetFutureInboundReqDto(_futureInboundReqDto *FutureInboundReqDto) error {
	r._futureInboundReqDto = _futureInboundReqDto
	r.Set("future_inbound_req_dto", _futureInboundReqDto)
	return nil
}

// GetFutureInboundReqDto FutureInboundReqDto Getter
func (r AlibabaAlihealthBcFutureStockInboundAPIRequest) GetFutureInboundReqDto() *FutureInboundReqDto {
	return r._futureInboundReqDto
}
