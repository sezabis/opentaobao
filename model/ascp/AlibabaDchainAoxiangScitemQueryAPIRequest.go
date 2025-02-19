package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemQueryAPIRequest 货品查询 API请求
// alibaba.dchain.aoxiang.scitem.query
//
// 货品查询
type AlibabaDchainAoxiangScitemQueryAPIRequest struct {
	model.Params
	// 货品查询入参
	_queryScitemRequest *QueryScItemRequest
}

// NewAlibabaDchainAoxiangScitemQueryRequest 初始化AlibabaDchainAoxiangScitemQueryAPIRequest对象
func NewAlibabaDchainAoxiangScitemQueryRequest() *AlibabaDchainAoxiangScitemQueryAPIRequest {
	return &AlibabaDchainAoxiangScitemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangScitemQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangScitemQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQueryScitemRequest is QueryScitemRequest Setter
// 货品查询入参
func (r *AlibabaDchainAoxiangScitemQueryAPIRequest) SetQueryScitemRequest(_queryScitemRequest *QueryScItemRequest) error {
	r._queryScitemRequest = _queryScitemRequest
	r.Set("query_scitem_request", _queryScitemRequest)
	return nil
}

// GetQueryScitemRequest QueryScitemRequest Getter
func (r AlibabaDchainAoxiangScitemQueryAPIRequest) GetQueryScitemRequest() *QueryScItemRequest {
	return r._queryScitemRequest
}
