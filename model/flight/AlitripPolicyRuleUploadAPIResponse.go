package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyRuleUploadAPIResponse 规则政策上传 API返回值
// alitrip.policy.rule.upload
//
// 上传特殊类型的单程/往返政策
type AlitripPolicyRuleUploadAPIResponse struct {
	model.CommonResponse
	AlitripPolicyRuleUploadAPIResponseModel
}

// AlitripPolicyRuleUploadAPIResponseModel is 规则政策上传 成功返回结果
type AlitripPolicyRuleUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_policy_rule_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlitripPolicyRuleUploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
