package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatAddCouponRecordAPIResponse 星河-记录用户微信优惠券领取记录 API返回值
// alitrip.merchant.galaxy.wechat.add.coupon.record
//
// 雅高小程序添加优惠券实例
type AlitripMerchantGalaxyWechatAddCouponRecordAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatAddCouponRecordAPIResponseModel
}

// AlitripMerchantGalaxyWechatAddCouponRecordAPIResponseModel is 星河-记录用户微信优惠券领取记录 成功返回结果
type AlitripMerchantGalaxyWechatAddCouponRecordAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_add_coupon_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlitripMerchantGalaxyWechatAddCouponRecordResponse `json:"result,omitempty" xml:"result,omitempty"`
}
