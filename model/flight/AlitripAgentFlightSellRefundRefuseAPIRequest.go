package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellRefundRefuseAPIRequest 销售退票拒绝 API请求
// alitrip.agent.flight.sell.refund.refuse
//
// 销售退票拒绝
type AlitripAgentFlightSellRefundRefuseAPIRequest struct {
	model.Params
	// 申请单号
	_applyId string
	// 拒绝原因
	_refuseReason string
	// 国内国际标识:1:国内,2:国际
	_domesticIntl int64
}

// NewAlitripAgentFlightSellRefundRefuseRequest 初始化AlitripAgentFlightSellRefundRefuseAPIRequest对象
func NewAlitripAgentFlightSellRefundRefuseRequest() *AlitripAgentFlightSellRefundRefuseAPIRequest {
	return &AlitripAgentFlightSellRefundRefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellRefundRefuseAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.refund.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellRefundRefuseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApplyId is ApplyId Setter
// 申请单号
func (r *AlitripAgentFlightSellRefundRefuseAPIRequest) SetApplyId(_applyId string) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlitripAgentFlightSellRefundRefuseAPIRequest) GetApplyId() string {
	return r._applyId
}

// SetRefuseReason is RefuseReason Setter
// 拒绝原因
func (r *AlitripAgentFlightSellRefundRefuseAPIRequest) SetRefuseReason(_refuseReason string) error {
	r._refuseReason = _refuseReason
	r.Set("refuse_reason", _refuseReason)
	return nil
}

// GetRefuseReason RefuseReason Getter
func (r AlitripAgentFlightSellRefundRefuseAPIRequest) GetRefuseReason() string {
	return r._refuseReason
}

// SetDomesticIntl is DomesticIntl Setter
// 国内国际标识:1:国内,2:国际
func (r *AlitripAgentFlightSellRefundRefuseAPIRequest) SetDomesticIntl(_domesticIntl int64) error {
	r._domesticIntl = _domesticIntl
	r.Set("domestic_intl", _domesticIntl)
	return nil
}

// GetDomesticIntl DomesticIntl Getter
func (r AlitripAgentFlightSellRefundRefuseAPIRequest) GetDomesticIntl() int64 {
	return r._domesticIntl
}
