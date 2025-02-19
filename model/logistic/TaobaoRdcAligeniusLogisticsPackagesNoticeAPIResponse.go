package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusLogisticsPackagesNoticeAPIResponse 物流多包裹通知 API返回值
// taobao.rdc.aligenius.logistics.packages.notice
//
// 订单发货之后，如果订单拆包、补发、赠品等场景，需要将多余包裹信息触达消费者
type TaobaoRdcAligeniusLogisticsPackagesNoticeAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusLogisticsPackagesNoticeAPIResponseModel
}

// TaobaoRdcAligeniusLogisticsPackagesNoticeAPIResponseModel is 物流多包裹通知 成功返回结果
type TaobaoRdcAligeniusLogisticsPackagesNoticeAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_logistics_packages_notice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoRdcAligeniusLogisticsPackagesNoticeResult `json:"result,omitempty" xml:"result,omitempty"`
}
