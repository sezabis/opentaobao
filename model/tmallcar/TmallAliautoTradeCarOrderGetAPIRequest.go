package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoTradeCarOrderGetAPIRequest 整车订单详情查询 API请求
// tmall.aliauto.trade.car.order.get
//
// 整车订单详情查询接口
type TmallAliautoTradeCarOrderGetAPIRequest struct {
	model.Params
	// 入参
	_query *SingleOrderDetailQuery
}

// NewTmallAliautoTradeCarOrderGetRequest 初始化TmallAliautoTradeCarOrderGetAPIRequest对象
func NewTmallAliautoTradeCarOrderGetRequest() *TmallAliautoTradeCarOrderGetAPIRequest {
	return &TmallAliautoTradeCarOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoTradeCarOrderGetAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.trade.car.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoTradeCarOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQuery is Query Setter
// 入参
func (r *TmallAliautoTradeCarOrderGetAPIRequest) SetQuery(_query *SingleOrderDetailQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallAliautoTradeCarOrderGetAPIRequest) GetQuery() *SingleOrderDetailQuery {
	return r._query
}
