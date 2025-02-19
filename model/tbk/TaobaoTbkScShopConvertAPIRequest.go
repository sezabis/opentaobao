package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScShopConvertAPIRequest 淘宝客-服务商-店铺链接转换 API请求
// taobao.tbk.sc.shop.convert
//
// 服务商使用。支持入参推广者对应的“推广位”和卖家id，获取对应的店铺推广链接。
type TaobaoTbkScShopConvertAPIRequest struct {
	model.Params
	// 需返回的字段列表
	_fields string
	// (该字段不开放)自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	_unid string
	// 卖家id串，用,分割，从taobao.tbk.shop.get接口获取user_id字段
	_userIds string
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneId int64
	// 链接形式-1：PC，2：无线。默认为１
	_platform int64
	// mm_xxx_xxx_xxx的第2段数字
	_siteId int64
}

// NewTaobaoTbkScShopConvertRequest 初始化TaobaoTbkScShopConvertAPIRequest对象
func NewTaobaoTbkScShopConvertRequest() *TaobaoTbkScShopConvertAPIRequest {
	return &TaobaoTbkScShopConvertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScShopConvertAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.shop.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScShopConvertAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需返回的字段列表
func (r *TaobaoTbkScShopConvertAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTbkScShopConvertAPIRequest) GetFields() string {
	return r._fields
}

// SetUnid is Unid Setter
// (该字段不开放)自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
func (r *TaobaoTbkScShopConvertAPIRequest) SetUnid(_unid string) error {
	r._unid = _unid
	r.Set("unid", _unid)
	return nil
}

// GetUnid Unid Getter
func (r TaobaoTbkScShopConvertAPIRequest) GetUnid() string {
	return r._unid
}

// SetUserIds is UserIds Setter
// 卖家id串，用,分割，从taobao.tbk.shop.get接口获取user_id字段
func (r *TaobaoTbkScShopConvertAPIRequest) SetUserIds(_userIds string) error {
	r._userIds = _userIds
	r.Set("user_ids", _userIds)
	return nil
}

// GetUserIds UserIds Getter
func (r TaobaoTbkScShopConvertAPIRequest) GetUserIds() string {
	return r._userIds
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaoTbkScShopConvertAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkScShopConvertAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

// SetPlatform is Platform Setter
// 链接形式-1：PC，2：无线。默认为１
func (r *TaobaoTbkScShopConvertAPIRequest) SetPlatform(_platform int64) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoTbkScShopConvertAPIRequest) GetPlatform() int64 {
	return r._platform
}

// SetSiteId is SiteId Setter
// mm_xxx_xxx_xxx的第2段数字
func (r *TaobaoTbkScShopConvertAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaoTbkScShopConvertAPIRequest) GetSiteId() int64 {
	return r._siteId
}
