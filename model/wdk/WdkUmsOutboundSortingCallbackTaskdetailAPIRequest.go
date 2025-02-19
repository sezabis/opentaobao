package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsOutboundSortingCallbackTaskdetailAPIRequest dps分货，明细回传 API请求
// wdk.ums.outbound.sorting.callback.taskdetail
//
// dps分货-分货明细回传
type WdkUmsOutboundSortingCallbackTaskdetailAPIRequest struct {
	model.Params
	// 参数
	_param0 *DpsCallBackSortMtopRequest
}

// NewWdkUmsOutboundSortingCallbackTaskdetailRequest 初始化WdkUmsOutboundSortingCallbackTaskdetailAPIRequest对象
func NewWdkUmsOutboundSortingCallbackTaskdetailRequest() *WdkUmsOutboundSortingCallbackTaskdetailAPIRequest {
	return &WdkUmsOutboundSortingCallbackTaskdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.callback.taskdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 参数
func (r *WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) SetParam0(_param0 *DpsCallBackSortMtopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) GetParam0() *DpsCallBackSortMtopRequest {
	return r._param0
}
