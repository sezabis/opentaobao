package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCouponTemplateSynAPIResponse 喵零券同步 API返回值
// tmall.nrt.coupon.template.syn
//
// 查询券模版
type TmallNrtCouponTemplateSynAPIResponse struct {
	model.CommonResponse
	TmallNrtCouponTemplateSynAPIResponseModel
}

// TmallNrtCouponTemplateSynAPIResponseModel is 喵零券同步 成功返回结果
type TmallNrtCouponTemplateSynAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_coupon_template_syn_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallNrtCouponTemplateSynResult `json:"result,omitempty" xml:"result,omitempty"`
}
