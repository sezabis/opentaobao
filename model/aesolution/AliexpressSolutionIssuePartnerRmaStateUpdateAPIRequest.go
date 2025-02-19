package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest aliexpress.solution.issue.partner.rma.state.update API请求
// aliexpress.solution.issue.partner.rma.state.update
//
// Receive changes in state updates for RMAs orders from after sales partners
type AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest struct {
	model.Params
	// RMA's order state update request
	_rmaStateUpdateRequest *RmaStateUpdateRequest
}

// NewAliexpressSolutionIssuePartnerRmaStateUpdateRequest 初始化AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest对象
func NewAliexpressSolutionIssuePartnerRmaStateUpdateRequest() *AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest {
	return &AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.issue.partner.rma.state.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRmaStateUpdateRequest is RmaStateUpdateRequest Setter
// RMA&#39;s order state update request
func (r *AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) SetRmaStateUpdateRequest(_rmaStateUpdateRequest *RmaStateUpdateRequest) error {
	r._rmaStateUpdateRequest = _rmaStateUpdateRequest
	r.Set("rma_state_update_request", _rmaStateUpdateRequest)
	return nil
}

// GetRmaStateUpdateRequest RmaStateUpdateRequest Getter
func (r AliexpressSolutionIssuePartnerRmaStateUpdateAPIRequest) GetRmaStateUpdateRequest() *RmaStateUpdateRequest {
	return r._rmaStateUpdateRequest
}
