package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionChangeQueryAPIResponse 改签航班询价 API返回值
// alitrip.btrip.flight.distribution.change.query
//
// 商旅机票改签航班询价
type AlitripBtripFlightDistributionChangeQueryAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionChangeQueryAPIResponseModel
}

// AlitripBtripFlightDistributionChangeQueryAPIResponseModel is 改签航班询价 成功返回结果
type AlitripBtripFlightDistributionChangeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_change_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
