package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemShiptimeUpdateAPIRequest 更新商品发货时间 API请求
// tmall.item.shiptime.update
//
// 增加更新删除商品/SKU发货时间(支持同一商品下的SKU同时批量更新)
// 1.
//     {
//         &#34;shipTimeType&#34;: 2,  ----相对发货时间（值为1则为绝对发货时间）
//         &#34;updateType&#34;: 0 ---更新SKU
//     },
//
//     按照指定SKU更新指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空
//
// 2.
//     {
//         &#34;shipTimeType&#34;: 0, -- 删除发货时间
//         &#34;updateType&#34;: 0 --更新SKU
//     },
//     按照指定SKU删除指定SKU的发货时间
//
// 3.
//
//     {
//         &#34;shipTimeType&#34;: 2,  ----相对发货时间（值为1则为绝对发货时间）
//         &#34;updateType&#34;: 1 ---更新商品
//     },
//
//     更新商品级发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间
//
// 4.
//     {
//         &#34;shipTimeType&#34;: 0, -- 删除发货时间
//         &#34;updateType&#34;: 1 --更新商品
//     },
//     删除商品级的发货时间
type TmallItemShiptimeUpdateAPIRequest struct {
	model.Params
	// 被更新SKU的发货时间，后台会根据三个子属性去查找匹配的sku，如果找到就默认对sku进行更新，当无匹配sku且更新类型针对sku，会报错。
	_skuShipTimes []UpdateSkuShipTime
	// 被更新发货时间（商品级）；格式和具体设置的发货时间格式相关。绝对发货时间填写yyyy-MM-dd;相对发货时间填写数字。发货时间必须在当前时间后三天。如果设置的绝对时间小于当前时间的三天后，会清除该商品的发货时间设置。如果是相对时间小于3，则会提示出错。如果shiptimeType为0，要清除商品上的发货时间，该字段可以填任意字符,也可以不填。
	_shipTime string
	// 商品ID
	_itemId int64
	// 批量更新商品/SKU发货时间的备选项
	_option *UpdateItemShipTimeOption
}

// NewTmallItemShiptimeUpdateRequest 初始化TmallItemShiptimeUpdateAPIRequest对象
func NewTmallItemShiptimeUpdateRequest() *TmallItemShiptimeUpdateAPIRequest {
	return &TmallItemShiptimeUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemShiptimeUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.shiptime.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemShiptimeUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSkuShipTimes is SkuShipTimes Setter
// 被更新SKU的发货时间，后台会根据三个子属性去查找匹配的sku，如果找到就默认对sku进行更新，当无匹配sku且更新类型针对sku，会报错。
func (r *TmallItemShiptimeUpdateAPIRequest) SetSkuShipTimes(_skuShipTimes []UpdateSkuShipTime) error {
	r._skuShipTimes = _skuShipTimes
	r.Set("sku_ship_times", _skuShipTimes)
	return nil
}

// GetSkuShipTimes SkuShipTimes Getter
func (r TmallItemShiptimeUpdateAPIRequest) GetSkuShipTimes() []UpdateSkuShipTime {
	return r._skuShipTimes
}

// SetShipTime is ShipTime Setter
// 被更新发货时间（商品级）；格式和具体设置的发货时间格式相关。绝对发货时间填写yyyy-MM-dd;相对发货时间填写数字。发货时间必须在当前时间后三天。如果设置的绝对时间小于当前时间的三天后，会清除该商品的发货时间设置。如果是相对时间小于3，则会提示出错。如果shiptimeType为0，要清除商品上的发货时间，该字段可以填任意字符,也可以不填。
func (r *TmallItemShiptimeUpdateAPIRequest) SetShipTime(_shipTime string) error {
	r._shipTime = _shipTime
	r.Set("ship_time", _shipTime)
	return nil
}

// GetShipTime ShipTime Getter
func (r TmallItemShiptimeUpdateAPIRequest) GetShipTime() string {
	return r._shipTime
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TmallItemShiptimeUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemShiptimeUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetOption is Option Setter
// 批量更新商品/SKU发货时间的备选项
func (r *TmallItemShiptimeUpdateAPIRequest) SetOption(_option *UpdateItemShipTimeOption) error {
	r._option = _option
	r.Set("option", _option)
	return nil
}

// GetOption Option Getter
func (r TmallItemShiptimeUpdateAPIRequest) GetOption() *UpdateItemShipTimeOption {
	return r._option
}
