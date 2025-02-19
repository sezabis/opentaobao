package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosIsvInventoryScrollqueryAPIRequest 滚动查询库存数据 API请求
// alibaba.mos.isv.inventory.scrollquery
//
// 按专柜滚动查询有库存商品
type AlibabaMosIsvInventoryScrollqueryAPIRequest struct {
	model.Params
	// 专柜ID
	_counterId string
	// 滚动查询ID号
	_scrollId string
}

// NewAlibabaMosIsvInventoryScrollqueryRequest 初始化AlibabaMosIsvInventoryScrollqueryAPIRequest对象
func NewAlibabaMosIsvInventoryScrollqueryRequest() *AlibabaMosIsvInventoryScrollqueryAPIRequest {
	return &AlibabaMosIsvInventoryScrollqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosIsvInventoryScrollqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.isv.inventory.scrollquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosIsvInventoryScrollqueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCounterId is CounterId Setter
// 专柜ID
func (r *AlibabaMosIsvInventoryScrollqueryAPIRequest) SetCounterId(_counterId string) error {
	r._counterId = _counterId
	r.Set("counter_id", _counterId)
	return nil
}

// GetCounterId CounterId Getter
func (r AlibabaMosIsvInventoryScrollqueryAPIRequest) GetCounterId() string {
	return r._counterId
}

// SetScrollId is ScrollId Setter
// 滚动查询ID号
func (r *AlibabaMosIsvInventoryScrollqueryAPIRequest) SetScrollId(_scrollId string) error {
	r._scrollId = _scrollId
	r.Set("scroll_id", _scrollId)
	return nil
}

// GetScrollId ScrollId Getter
func (r AlibabaMosIsvInventoryScrollqueryAPIRequest) GetScrollId() string {
	return r._scrollId
}
