package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallCarOrderQueryAPIResponse 天猫汽车整车订单查询 API返回值
// tmall.car.order.query
//
// 天猫汽车商家通过该接口查看整车订单信息
type TmallCarOrderQueryAPIResponse struct {
	model.CommonResponse
	TmallCarOrderQueryAPIResponseModel
}

// TmallCarOrderQueryAPIResponseModel is 天猫汽车整车订单查询 成功返回结果
type TmallCarOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页结果
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
