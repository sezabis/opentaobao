package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomOrderCheckorderinfoAPIResponse 金融购机订单信息校验 API返回值
// alibaba.alicom.order.checkorderinfo
//
// 金融购机订单信息校验
type AlibabaAlicomOrderCheckorderinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlicomOrderCheckorderinfoAPIResponseModel
}

// AlibabaAlicomOrderCheckorderinfoAPIResponseModel is 金融购机订单信息校验 成功返回结果
type AlibabaAlicomOrderCheckorderinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_order_checkorderinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
