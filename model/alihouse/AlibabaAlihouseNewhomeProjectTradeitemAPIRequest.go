package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectTradeitemAPIRequest 新房品同步 API请求
// alibaba.alihouse.newhome.project.tradeitem
//
// 新房品同步
type AlibabaAlihouseNewhomeProjectTradeitemAPIRequest struct {
	model.Params
	// 请求对象
	_goodsDto *GoodsDto
}

// NewAlibabaAlihouseNewhomeProjectTradeitemRequest 初始化AlibabaAlihouseNewhomeProjectTradeitemAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectTradeitemRequest() *AlibabaAlihouseNewhomeProjectTradeitemAPIRequest {
	return &AlibabaAlihouseNewhomeProjectTradeitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectTradeitemAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.tradeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectTradeitemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGoodsDto is GoodsDto Setter
// 请求对象
func (r *AlibabaAlihouseNewhomeProjectTradeitemAPIRequest) SetGoodsDto(_goodsDto *GoodsDto) error {
	r._goodsDto = _goodsDto
	r.Set("goods_dto", _goodsDto)
	return nil
}

// GetGoodsDto GoodsDto Getter
func (r AlibabaAlihouseNewhomeProjectTradeitemAPIRequest) GetGoodsDto() *GoodsDto {
	return r._goodsDto
}
