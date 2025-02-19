package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScOrderDetailsTemporaryGetAPIResponse 淘宝客-服务商-所有订单查询（临时接口） API返回值
// taobao.tbk.sc.order.details.temporary.get
//
// 服务商使用。可通过该接口查询推广者账号下对应的推广订单明细。
type TaobaoTbkScOrderDetailsTemporaryGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScOrderDetailsTemporaryGetAPIResponseModel
}

// TaobaoTbkScOrderDetailsTemporaryGetAPIResponseModel is 淘宝客-服务商-所有订单查询（临时接口） 成功返回结果
type TaobaoTbkScOrderDetailsTemporaryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_order_details_temporary_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// PublisherOrderDto
	Data *OrderPage `json:"data,omitempty" xml:"data,omitempty"`
}
