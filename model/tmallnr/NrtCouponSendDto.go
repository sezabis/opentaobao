package tmallnr

// NrtCouponSendDto 结构体
type NrtCouponSendDto struct {
	// 加密后淘系ID
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 业务code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// {&#34;key&#34;,&#34;value&#34;}
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 幂等ID
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 券code
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 券类型，247：品类券，276：门店券，357：门店通用券
	CouponType int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
}
