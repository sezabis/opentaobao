package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatCardParmQueryAPIResponse 微信会员卡添加 API返回值
// alitrip.merchant.galaxy.wechat.card.parm.query
//
// 微信会员卡添加参数获取
type AlitripMerchantGalaxyWechatCardParmQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatCardParmQueryAPIResponseModel
}

// AlitripMerchantGalaxyWechatCardParmQueryAPIResponseModel is 微信会员卡添加 成功返回结果
type AlitripMerchantGalaxyWechatCardParmQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_card_parm_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlitripMerchantGalaxyWechatCardParmQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
