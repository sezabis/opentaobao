package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantMaDelayAPIRequest
凭证延期 API请求
taobao.eticket.merchant.ma.delay

订单延期 */
type TaobaoEticketMerchantMaDelayAPIRequest struct {
	model.Params
	// 业务类型
	_bizType int64
	// 延期时间
	_endDate string
	// 码
	_code string
	// 订单号
	_outerId string
	// 扩展
	_attributeMap string
	// 请求ID，调用方保证惟一
	_requestId string
}

// NewTaobaoEticketMerchantMaDelayRequest 初始化TaobaoEticketMerchantMaDelayAPIRequest对象
func NewTaobaoEticketMerchantMaDelayRequest() *TaobaoEticketMerchantMaDelayAPIRequest {
	return &TaobaoEticketMerchantMaDelayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaDelayAPIRequest) GetApiMethodName() string {
	return "taobao.eticket.merchant.ma.delay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaDelayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaDelayAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// Get BizType Getter
func (r TaobaoEticketMerchantMaDelayAPIRequest) GetBizType() int64 {
	return r._bizType
}

// Set is EndDate Setter
// 延期时间
func (r *TaobaoEticketMerchantMaDelayAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r TaobaoEticketMerchantMaDelayAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is Code Setter
// 码
func (r *TaobaoEticketMerchantMaDelayAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r TaobaoEticketMerchantMaDelayAPIRequest) GetCode() string {
	return r._code
}

// Set is OuterId Setter
// 订单号
func (r *TaobaoEticketMerchantMaDelayAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TaobaoEticketMerchantMaDelayAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is AttributeMap Setter
// 扩展
func (r *TaobaoEticketMerchantMaDelayAPIRequest) SetAttributeMap(_attributeMap string) error {
	r._attributeMap = _attributeMap
	r.Set("attribute_map", _attributeMap)
	return nil
}

// Get AttributeMap Getter
func (r TaobaoEticketMerchantMaDelayAPIRequest) GetAttributeMap() string {
	return r._attributeMap
}

// Set is RequestId Setter
// 请求ID，调用方保证惟一
func (r *TaobaoEticketMerchantMaDelayAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// Get RequestId Getter
func (r TaobaoEticketMerchantMaDelayAPIRequest) GetRequestId() string {
	return r._requestId
}
