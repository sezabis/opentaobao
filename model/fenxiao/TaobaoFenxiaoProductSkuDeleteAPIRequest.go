package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductSkuDeleteAPIRequest
产品SKU删除接口 API请求
taobao.fenxiao.product.sku.delete

根据sku properties删除sku数据 */
type TaobaoFenxiaoProductSkuDeleteAPIRequest struct {
	model.Params
	// 产品id
	_productId int64
	// sku属性
	_properties string
}

// NewTaobaoFenxiaoProductSkuDeleteRequest 初始化TaobaoFenxiaoProductSkuDeleteAPIRequest对象
func NewTaobaoFenxiaoProductSkuDeleteRequest() *TaobaoFenxiaoProductSkuDeleteAPIRequest {
	return &TaobaoFenxiaoProductSkuDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductSkuDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkuDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProductId Setter
// 产品id
func (r *TaobaoFenxiaoProductSkuDeleteAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r TaobaoFenxiaoProductSkuDeleteAPIRequest) GetProductId() int64 {
	return r._productId
}

// Set is Properties Setter
// sku属性
func (r *TaobaoFenxiaoProductSkuDeleteAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// Get Properties Getter
func (r TaobaoFenxiaoProductSkuDeleteAPIRequest) GetProperties() string {
	return r._properties
}
