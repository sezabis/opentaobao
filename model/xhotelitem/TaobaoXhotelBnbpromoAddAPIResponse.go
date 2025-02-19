package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoAddAPIResponse 自促活动申请接口 API返回值
// taobao.xhotel.bnbpromo.add
//
// 自促活动申请接口
type TaobaoXhotelBnbpromoAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbpromoAddAPIResponseModel
}

// TaobaoXhotelBnbpromoAddAPIResponseModel is 自促活动申请接口 成功返回结果
type TaobaoXhotelBnbpromoAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbpromo_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销添加返回对象
	Module *PromoCode `json:"module,omitempty" xml:"module,omitempty"`
}
