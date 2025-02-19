package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCouponTemplateSynAPIRequest 喵零券同步 API请求
// tmall.nrt.coupon.template.syn
//
// 查询券模版
type TmallNrtCouponTemplateSynAPIRequest struct {
	model.Params
	// 业务编码
	_bizCode string
	// 券模版id集合
	_couponTemplateId string
	// 券类型
	_couponType int64
}

// NewTmallNrtCouponTemplateSynRequest 初始化TmallNrtCouponTemplateSynAPIRequest对象
func NewTmallNrtCouponTemplateSynRequest() *TmallNrtCouponTemplateSynAPIRequest {
	return &TmallNrtCouponTemplateSynAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtCouponTemplateSynAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.coupon.template.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtCouponTemplateSynAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizCode is BizCode Setter
// 业务编码
func (r *TmallNrtCouponTemplateSynAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallNrtCouponTemplateSynAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetCouponTemplateId is CouponTemplateId Setter
// 券模版id集合
func (r *TmallNrtCouponTemplateSynAPIRequest) SetCouponTemplateId(_couponTemplateId string) error {
	r._couponTemplateId = _couponTemplateId
	r.Set("coupon_template_id", _couponTemplateId)
	return nil
}

// GetCouponTemplateId CouponTemplateId Getter
func (r TmallNrtCouponTemplateSynAPIRequest) GetCouponTemplateId() string {
	return r._couponTemplateId
}

// SetCouponType is CouponType Setter
// 券类型
func (r *TmallNrtCouponTemplateSynAPIRequest) SetCouponType(_couponType int64) error {
	r._couponType = _couponType
	r.Set("coupon_type", _couponType)
	return nil
}

// GetCouponType CouponType Getter
func (r TmallNrtCouponTemplateSynAPIRequest) GetCouponType() int64 {
	return r._couponType
}
