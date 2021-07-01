package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMobilePromotionCouponApplyAPIRequest
优惠券领取(手淘专用) API请求
taobao.mobile.promotion.coupon.apply

优惠券领取 */
type TaobaoMobilePromotionCouponApplyAPIRequest struct {
	model.Params
	// 请求唯一id，问题排查
	_traceId string
	// 传播id
	_spreadId int64
	// 广播id
	_feedId string
	// 三方活动id
	_bizId string
}

// NewTaobaoMobilePromotionCouponApplyRequest 初始化TaobaoMobilePromotionCouponApplyAPIRequest对象
func NewTaobaoMobilePromotionCouponApplyRequest() *TaobaoMobilePromotionCouponApplyAPIRequest {
	return &TaobaoMobilePromotionCouponApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetApiMethodName() string {
	return "taobao.mobile.promotion.coupon.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TraceId Setter
// 请求唯一id，问题排查
func (r *TaobaoMobilePromotionCouponApplyAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// Get TraceId Getter
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetTraceId() string {
	return r._traceId
}

// Set is SpreadId Setter
// 传播id
func (r *TaobaoMobilePromotionCouponApplyAPIRequest) SetSpreadId(_spreadId int64) error {
	r._spreadId = _spreadId
	r.Set("spread_id", _spreadId)
	return nil
}

// Get SpreadId Getter
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetSpreadId() int64 {
	return r._spreadId
}

// Set is FeedId Setter
// 广播id
func (r *TaobaoMobilePromotionCouponApplyAPIRequest) SetFeedId(_feedId string) error {
	r._feedId = _feedId
	r.Set("feed_id", _feedId)
	return nil
}

// Get FeedId Getter
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetFeedId() string {
	return r._feedId
}

// Set is BizId Setter
// 三方活动id
func (r *TaobaoMobilePromotionCouponApplyAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// Get BizId Getter
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetBizId() string {
	return r._bizId
}
