package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemsCustomGetAPIRequest 根据外部ID取商品 API请求
// taobao.items.custom.get
//
// 跟据卖家设定的商品外部id获取商品，只能获取授权卖家的商品
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoItemsCustomGetAPIRequest struct {
	model.Params
	// 商品的外部商品ID，支持批量，最多不超过40个。
	_outerId string
	// 需返回的字段列表，参考：Item商品结构体说明，其中barcode、sku.barcode等条形码字段暂不支持；多个字段之间用“,”分隔。
	_fields string
}

// NewTaobaoItemsCustomGetRequest 初始化TaobaoItemsCustomGetAPIRequest对象
func NewTaobaoItemsCustomGetRequest() *TaobaoItemsCustomGetAPIRequest {
	return &TaobaoItemsCustomGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemsCustomGetAPIRequest) GetApiMethodName() string {
	return "taobao.items.custom.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemsCustomGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterId is OuterId Setter
// 商品的外部商品ID，支持批量，最多不超过40个。
func (r *TaobaoItemsCustomGetAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoItemsCustomGetAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetFields is Fields Setter
// 需返回的字段列表，参考：Item商品结构体说明，其中barcode、sku.barcode等条形码字段暂不支持；多个字段之间用“,”分隔。
func (r *TaobaoItemsCustomGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoItemsCustomGetAPIRequest) GetFields() string {
	return r._fields
}
