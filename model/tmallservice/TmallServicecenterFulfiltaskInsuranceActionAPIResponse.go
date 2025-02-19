package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterFulfiltaskInsuranceActionAPIResponse 供应链保险链路动作 API返回值
// tmall.servicecenter.fulfiltask.insurance.action
//
// 服务供应链履约链路 保险类业务履约接口
type TmallServicecenterFulfiltaskInsuranceActionAPIResponse struct {
	model.CommonResponse
	TmallServicecenterFulfiltaskInsuranceActionAPIResponseModel
}

// TmallServicecenterFulfiltaskInsuranceActionAPIResponseModel is 供应链保险链路动作 成功返回结果
type TmallServicecenterFulfiltaskInsuranceActionAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_fulfiltask_insurance_action_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallServicecenterFulfiltaskInsuranceActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
