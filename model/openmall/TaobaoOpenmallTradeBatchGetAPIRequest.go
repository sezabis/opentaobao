package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeBatchGetAPIRequest
批量获取openmall订单 API请求
taobao.openmall.trade.batch.get

批量获取openmall订单
注意：该接口数据存在延迟，实时数据请通过taobao.openmall.trade.get获取 */
type TaobaoOpenmallTradeBatchGetAPIRequest struct {
	model.Params
	// 查询范围结束时间，闭区间
	_endCreated string
	// 查询页码，从1开始
	_pageIndex int64
	// 页面大小，不超过100
	_pageSize int64
	// 渠道商Nick
	_distributor string
	// 查询范围开始时间，闭区间
	_startCreated string
}

// NewTaobaoOpenmallTradeBatchGetRequest 初始化TaobaoOpenmallTradeBatchGetAPIRequest对象
func NewTaobaoOpenmallTradeBatchGetRequest() *TaobaoOpenmallTradeBatchGetAPIRequest {
	return &TaobaoOpenmallTradeBatchGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.batch.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EndCreated Setter
// 查询范围结束时间，闭区间
func (r *TaobaoOpenmallTradeBatchGetAPIRequest) SetEndCreated(_endCreated string) error {
	r._endCreated = _endCreated
	r.Set("end_created", _endCreated)
	return nil
}

// Get EndCreated Getter
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetEndCreated() string {
	return r._endCreated
}

// Set is PageIndex Setter
// 查询页码，从1开始
func (r *TaobaoOpenmallTradeBatchGetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// Get PageIndex Getter
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// Set is PageSize Setter
// 页面大小，不超过100
func (r *TaobaoOpenmallTradeBatchGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is Distributor Setter
// 渠道商Nick
func (r *TaobaoOpenmallTradeBatchGetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// Get Distributor Getter
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetDistributor() string {
	return r._distributor
}

// Set is StartCreated Setter
// 查询范围开始时间，闭区间
func (r *TaobaoOpenmallTradeBatchGetAPIRequest) SetStartCreated(_startCreated string) error {
	r._startCreated = _startCreated
	r.Set("start_created", _startCreated)
	return nil
}

// Get StartCreated Getter
func (r TaobaoOpenmallTradeBatchGetAPIRequest) GetStartCreated() string {
	return r._startCreated
}
