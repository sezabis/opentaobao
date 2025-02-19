package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderFillAPIResponse 填单页接口 API返回值
// alitrip.merchant.galaxy.order.fill
//
// 进入填单页时调用此接口，返回填单所需展示基础信息
type AlitripMerchantGalaxyOrderFillAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyOrderFillAPIResponseModel
}

// AlitripMerchantGalaxyOrderFillAPIResponseModel is 填单页接口 成功返回结果
type AlitripMerchantGalaxyOrderFillAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_fill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyOrderFillResponse `json:"result,omitempty" xml:"result,omitempty"`
}
