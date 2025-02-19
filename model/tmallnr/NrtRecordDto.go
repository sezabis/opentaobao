package tmallnr

// NrtRecordDto 结构体
type NrtRecordDto struct {
	// 券信息集合
	CouponList []NrtCrmBenefitDto `json:"coupon_list,omitempty" xml:"coupon_list>nrt_crm_benefit_dto,omitempty"`
	// 有效结束时间
	CertificateEndTime string `json:"certificate_end_time,omitempty" xml:"certificate_end_time,omitempty"`
	// 有效开始时间
	CertificateStartTime string `json:"certificate_start_time,omitempty" xml:"certificate_start_time,omitempty"`
	// 电子凭证编码
	CertificateCode string `json:"certificate_code,omitempty" xml:"certificate_code,omitempty"`
	// 加密后的会员ID
	OpenIdStr string `json:"open_id_str,omitempty" xml:"open_id_str,omitempty"`
	// 会员openId-废弃
	OpenId int64 `json:"open_id,omitempty" xml:"open_id,omitempty"`
}
