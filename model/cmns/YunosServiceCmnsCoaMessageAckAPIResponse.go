package cmns

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaMessageAckAPIResponse 消息回执查询 API返回值
// yunos.service.cmns.coa.message.ack
//
// 第三方应用开发者调用此接口查询设备是否收到消息，只能查询此appKey床发的消息
type YunosServiceCmnsCoaMessageAckAPIResponse struct {
	model.CommonResponse
	YunosServiceCmnsCoaMessageAckAPIResponseModel
}

// YunosServiceCmnsCoaMessageAckAPIResponseModel is 消息回执查询 成功返回结果
type YunosServiceCmnsCoaMessageAckAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_service_cmns_coa_message_ack_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口出错提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 接口调用成功
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 接口调用成功返回信息&lt;br/&gt;0:未到达 1：已到达
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
}
