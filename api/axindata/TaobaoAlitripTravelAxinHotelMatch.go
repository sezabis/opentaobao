package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelMatch 酒店匹配接口-阿信
// taobao.alitrip.travel.axin.hotel.match
//
// 酒店匹配接口-阿信
func TaobaoAlitripTravelAxinHotelMatch(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelMatchAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinHotelMatchAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinHotelMatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
