package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectQueryAPIRequest 查询楼盘相关信息 API请求
// alibaba.alihouse.newhome.project.query
//
// 根据outerid查询楼盘相关信息
type AlibabaAlihouseNewhomeProjectQueryAPIRequest struct {
	model.Params
	// 外部楼盘/小区id
	_outerId string
	// 商品id
	_itemId string
}

// NewAlibabaAlihouseNewhomeProjectQueryRequest 初始化AlibabaAlihouseNewhomeProjectQueryAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectQueryRequest() *AlibabaAlihouseNewhomeProjectQueryAPIRequest {
	return &AlibabaAlihouseNewhomeProjectQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterId is OuterId Setter
// 外部楼盘/小区id
func (r *AlibabaAlihouseNewhomeProjectQueryAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeProjectQueryAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlibabaAlihouseNewhomeProjectQueryAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaAlihouseNewhomeProjectQueryAPIRequest) GetItemId() string {
	return r._itemId
}
