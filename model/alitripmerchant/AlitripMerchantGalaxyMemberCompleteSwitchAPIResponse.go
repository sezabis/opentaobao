package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyMemberCompleteSwitchAPIResponse 会员切换模式接口 API返回值
// alitrip.merchant.galaxy.member.complete.switch
//
// 小程序老用户调用德比接口进行会员切换
type AlitripMerchantGalaxyMemberCompleteSwitchAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberCompleteSwitchAPIResponseModel
}

// AlitripMerchantGalaxyMemberCompleteSwitchAPIResponseModel is 会员切换模式接口 成功返回结果
type AlitripMerchantGalaxyMemberCompleteSwitchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_complete_switch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *AlitripMerchantGalaxyMemberCompleteSwitchResponse `json:"result,omitempty" xml:"result,omitempty"`
}
