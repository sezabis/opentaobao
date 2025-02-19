package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoTradeCarOrderGet 整车订单详情查询
// tmall.aliauto.trade.car.order.get
//
// 整车订单详情查询接口
func TmallAliautoTradeCarOrderGet(clt *core.SDKClient, req *tmallcar.TmallAliautoTradeCarOrderGetAPIRequest, session string) (*tmallcar.TmallAliautoTradeCarOrderGetAPIResponse, error) {
	var resp tmallcar.TmallAliautoTradeCarOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
