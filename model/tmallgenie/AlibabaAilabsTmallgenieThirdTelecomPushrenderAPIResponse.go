package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse 电信-推送拉起设备应用 API返回值
// alibaba.ailabs.tmallgenie.third.telecom.pushrender
//
// 电信-推送拉起设备应用
type AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponseModel
}

// AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponseModel is 电信-推送拉起设备应用 成功返回结果
type AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_third_telecom_pushrender_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *AlibabaAilabsTmallgenieThirdTelecomPushrenderResult `json:"result,omitempty" xml:"result,omitempty"`
}
