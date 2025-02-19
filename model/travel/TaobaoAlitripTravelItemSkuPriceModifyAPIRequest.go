package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemSkuPriceModifyAPIRequest 【API3.0】日期级别日历价格库存修改，增量维护 API请求
// taobao.alitrip.travel.item.sku.price.modify
//
// 【API3.0】日期级别日历价格库存增量维护
type TaobaoAlitripTravelItemSkuPriceModifyAPIRequest struct {
	model.Params
	// 商品日历价格库存套餐
	_skus []PontusTravelItemSkuInfo
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品价库变更类型，0=价格库存均变更，1=价格变更，2=库存变更
	_modifyType int64
}

// NewTaobaoAlitripTravelItemSkuPriceModifyRequest 初始化TaobaoAlitripTravelItemSkuPriceModifyAPIRequest对象
func NewTaobaoAlitripTravelItemSkuPriceModifyRequest() *TaobaoAlitripTravelItemSkuPriceModifyAPIRequest {
	return &TaobaoAlitripTravelItemSkuPriceModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.sku.price.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSkus is Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) SetSkus(_skus []PontusTravelItemSkuInfo) error {
	r._skus = _skus
	r.Set("skus", _skus)
	return nil
}

// GetSkus Skus Getter
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetSkus() []PontusTravelItemSkuInfo {
	return r._skus
}

// SetOutProductId is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetModifyType is ModifyType Setter
// 商品价库变更类型，0=价格库存均变更，1=价格变更，2=库存变更
func (r *TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) SetModifyType(_modifyType int64) error {
	r._modifyType = _modifyType
	r.Set("modify_type", _modifyType)
	return nil
}

// GetModifyType ModifyType Getter
func (r TaobaoAlitripTravelItemSkuPriceModifyAPIRequest) GetModifyType() int64 {
	return r._modifyType
}
