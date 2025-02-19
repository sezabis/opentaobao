package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelDetailQuery 阿信酒店分销-标准酒店详情查询
// taobao.alitrip.travel.axin.hotel.detail.query
//
// 标准酒店详情查询
func TaobaoAlitripTravelAxinHotelDetailQuery(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinHotelDetailQueryAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinHotelDetailQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
