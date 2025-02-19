package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest 工人取消请假 API请求
// alibaba.ssc.supplyplatform.serviceworker.cancelleave
//
// 工人取消请假
type AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest struct {
	model.Params
	// 时间段
	_leaveBeginAndEndList []string
	// 身份证号
	_identityNumber string
}

// NewAlibabaSscSupplyplatformServiceworkerCancelleaveRequest 初始化AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest对象
func NewAlibabaSscSupplyplatformServiceworkerCancelleaveRequest() *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest {
	return &AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceworker.cancelleave"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLeaveBeginAndEndList is LeaveBeginAndEndList Setter
// 时间段
func (r *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) SetLeaveBeginAndEndList(_leaveBeginAndEndList []string) error {
	r._leaveBeginAndEndList = _leaveBeginAndEndList
	r.Set("leave_begin_and_end_list", _leaveBeginAndEndList)
	return nil
}

// GetLeaveBeginAndEndList LeaveBeginAndEndList Getter
func (r AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) GetLeaveBeginAndEndList() []string {
	return r._leaveBeginAndEndList
}

// SetIdentityNumber is IdentityNumber Setter
// 身份证号
func (r *AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) SetIdentityNumber(_identityNumber string) error {
	r._identityNumber = _identityNumber
	r.Set("identity_number", _identityNumber)
	return nil
}

// GetIdentityNumber IdentityNumber Getter
func (r AlibabaSscSupplyplatformServiceworkerCancelleaveAPIRequest) GetIdentityNumber() string {
	return r._identityNumber
}
