package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoFulfillmentDeliverySynAPIResponse 交付状态及物流信息同步 API返回值
// tmall.aliauto.fulfillment.delivery.syn
//
// 交付状态及物流信息同步
type TmallAliautoFulfillmentDeliverySynAPIResponse struct {
	model.CommonResponse
	TmallAliautoFulfillmentDeliverySynAPIResponseModel
}

// TmallAliautoFulfillmentDeliverySynAPIResponseModel is 交付状态及物流信息同步 成功返回结果
type TmallAliautoFulfillmentDeliverySynAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_fulfillment_delivery_syn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}
