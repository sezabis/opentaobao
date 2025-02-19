package perfect

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse 装箱接单 API返回值
// alibaba.tcwms.outbound.load.container.receive
//
// 装箱接单
type AlibabaTcwmsOutboundLoadContainerReceiveAPIResponse struct {
	model.CommonResponse
	AlibabaTcwmsOutboundLoadContainerReceiveAPIResponseModel
}

// AlibabaTcwmsOutboundLoadContainerReceiveAPIResponseModel is 装箱接单 成功返回结果
type AlibabaTcwmsOutboundLoadContainerReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcwms_outbound_load_container_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数
	Result *LoadReceiveResponse `json:"result,omitempty" xml:"result,omitempty"`
}
