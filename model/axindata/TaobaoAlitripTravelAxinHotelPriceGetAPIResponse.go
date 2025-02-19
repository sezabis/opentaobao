package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelPriceGetAPIResponse 酒店报价服务-阿信 API返回值
// taobao.alitrip.travel.axin.hotel.price.get
//
// 酒店报价查询服务
type TaobaoAlitripTravelAxinHotelPriceGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinHotelPriceGetAPIResponseModel
}

// TaobaoAlitripTravelAxinHotelPriceGetAPIResponseModel is 酒店报价服务-阿信 成功返回结果
type TaobaoAlitripTravelAxinHotelPriceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_hotel_price_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripTravelAxinHotelPriceGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
