package alitripmerchant

// CouponInfo 结构体
type CouponInfo struct {
	// 使用描述
	DetailDesc string `json:"detail_desc,omitempty" xml:"detail_desc,omitempty"`
	// 优惠券名称
	CouponName string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	// 到期日期
	ExpiredTimeMin string `json:"expired_time_min,omitempty" xml:"expired_time_min,omitempty"`
	// 面额(减免额度) 元
	DiscountAmount string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 使用条件描述
	ConditionDesc string `json:"condition_desc,omitempty" xml:"condition_desc,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}
