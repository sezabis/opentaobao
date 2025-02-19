package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyVoucherQueryListAPIResponse 查询代金券列表 API返回值
// alitrip.merchant.galaxy.voucher.query.list
//
// 查询代金券列表
type AlitripMerchantGalaxyVoucherQueryListAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyVoucherQueryListAPIResponseModel
}

// AlitripMerchantGalaxyVoucherQueryListAPIResponseModel is 查询代金券列表 成功返回结果
type AlitripMerchantGalaxyVoucherQueryListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_voucher_query_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyVoucherQueryListResponse `json:"result,omitempty" xml:"result,omitempty"`
}
