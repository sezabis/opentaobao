package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderFailAPIResponse 出票失败 API返回值
// taobao.train.agent.order.fail
//
// 出票失败
type TaobaoTrainAgentOrderFailAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderFailAPIResponseModel
}

// TaobaoTrainAgentOrderFailAPIResponseModel is 出票失败 成功返回结果
type TaobaoTrainAgentOrderFailAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_fail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}
