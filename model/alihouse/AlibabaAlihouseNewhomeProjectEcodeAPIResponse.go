package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectEcodeAPIResponse 楼盘e码更新 API返回值
// alibaba.alihouse.newhome.project.ecode
//
// 更新楼盘ecode
type AlibabaAlihouseNewhomeProjectEcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectEcodeAPIResponseModel
}

// AlibabaAlihouseNewhomeProjectEcodeAPIResponseModel is 楼盘e码更新 成功返回结果
type AlibabaAlihouseNewhomeProjectEcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_ecode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectEcodeResult `json:"result,omitempty" xml:"result,omitempty"`
}
