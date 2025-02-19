package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtSmsCodeSendAPIResponse 喵零发送短信 API返回值
// tmall.nrt.sms.code.send
//
// 喵零发送短信
type TmallNrtSmsCodeSendAPIResponse struct {
	model.CommonResponse
	TmallNrtSmsCodeSendAPIResponseModel
}

// TmallNrtSmsCodeSendAPIResponseModel is 喵零发送短信 成功返回结果
type TmallNrtSmsCodeSendAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_sms_code_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 请求成功与否
	Succ string `json:"succ,omitempty" xml:"succ,omitempty"`
	// 系统异常
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据对象
	Data *NrtSmsDto `json:"data,omitempty" xml:"data,omitempty"`
}
