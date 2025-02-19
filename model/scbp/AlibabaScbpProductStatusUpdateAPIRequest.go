package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpProductStatusUpdateAPIRequest 修改P4P产品推广状态 API请求
// alibaba.scbp.product.status.update
//
// 修改P4P产品推广状态
type AlibabaScbpProductStatusUpdateAPIRequest struct {
	model.Params
	// 产品ID列表
	_productIdList []string
	// enabled:开启,disabled:暂停
	_status string
}

// NewAlibabaScbpProductStatusUpdateRequest 初始化AlibabaScbpProductStatusUpdateAPIRequest对象
func NewAlibabaScbpProductStatusUpdateRequest() *AlibabaScbpProductStatusUpdateAPIRequest {
	return &AlibabaScbpProductStatusUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.product.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductStatusUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProductIdList is ProductIdList Setter
// 产品ID列表
func (r *AlibabaScbpProductStatusUpdateAPIRequest) SetProductIdList(_productIdList []string) error {
	r._productIdList = _productIdList
	r.Set("product_id_list", _productIdList)
	return nil
}

// GetProductIdList ProductIdList Getter
func (r AlibabaScbpProductStatusUpdateAPIRequest) GetProductIdList() []string {
	return r._productIdList
}

// SetStatus is Status Setter
// enabled:开启,disabled:暂停
func (r *AlibabaScbpProductStatusUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaScbpProductStatusUpdateAPIRequest) GetStatus() string {
	return r._status
}
