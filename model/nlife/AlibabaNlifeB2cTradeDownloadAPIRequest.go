package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cTradeDownloadAPIRequest
b2c下载订单 API请求
alibaba.nlife.b2c.trade.download

下载零售商在零售+平台创建的订单 */
type AlibabaNlifeB2cTradeDownloadAPIRequest struct {
	model.Params
	// 页码
	_pageNo int64
	// 分页大小
	_pageSize int64
	// 零售门店在零售+平台对应的ID
	_storeId string
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
}

// NewAlibabaNlifeB2cTradeDownloadRequest 初始化AlibabaNlifeB2cTradeDownloadAPIRequest对象
func NewAlibabaNlifeB2cTradeDownloadRequest() *AlibabaNlifeB2cTradeDownloadAPIRequest {
	return &AlibabaNlifeB2cTradeDownloadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradeDownloadAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.b2c.trade.download"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradeDownloadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PageNo Setter
// 页码
func (r *AlibabaNlifeB2cTradeDownloadAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r AlibabaNlifeB2cTradeDownloadAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 分页大小
func (r *AlibabaNlifeB2cTradeDownloadAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaNlifeB2cTradeDownloadAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is StoreId Setter
// 零售门店在零售+平台对应的ID
func (r *AlibabaNlifeB2cTradeDownloadAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaNlifeB2cTradeDownloadAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is StartDate Setter
// 开始时间
func (r *AlibabaNlifeB2cTradeDownloadAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r AlibabaNlifeB2cTradeDownloadAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is EndDate Setter
// 结束时间
func (r *AlibabaNlifeB2cTradeDownloadAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaNlifeB2cTradeDownloadAPIRequest) GetEndDate() string {
	return r._endDate
}
