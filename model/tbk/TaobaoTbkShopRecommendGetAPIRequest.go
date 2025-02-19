package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkShopRecommendGetAPIRequest 淘宝客-公用-店铺关联推荐 API请求
// taobao.tbk.shop.recommend.get
//
// 入参卖家id，可推荐与此店铺相关联的相关店铺。
type TaobaoTbkShopRecommendGetAPIRequest struct {
	model.Params
	// 需返回的字段列表
	_fields string
	// 返回数量，默认20，最大值40
	_count int64
	// 链接形式：1：PC，2：无线，默认：１
	_platform int64
	// 卖家Id
	_userId int64
}

// NewTaobaoTbkShopRecommendGetRequest 初始化TaobaoTbkShopRecommendGetAPIRequest对象
func NewTaobaoTbkShopRecommendGetRequest() *TaobaoTbkShopRecommendGetAPIRequest {
	return &TaobaoTbkShopRecommendGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkShopRecommendGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.shop.recommend.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkShopRecommendGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需返回的字段列表
func (r *TaobaoTbkShopRecommendGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTbkShopRecommendGetAPIRequest) GetFields() string {
	return r._fields
}

// SetCount is Count Setter
// 返回数量，默认20，最大值40
func (r *TaobaoTbkShopRecommendGetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r TaobaoTbkShopRecommendGetAPIRequest) GetCount() int64 {
	return r._count
}

// SetPlatform is Platform Setter
// 链接形式：1：PC，2：无线，默认：１
func (r *TaobaoTbkShopRecommendGetAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoTbkShopRecommendGetAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetUserId is UserId Setter
// 卖家Id
func (r *TaobaoTbkShopRecommendGetAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoTbkShopRecommendGetAPIRequest) GetUserId() int64 {
	return r._userId
}
