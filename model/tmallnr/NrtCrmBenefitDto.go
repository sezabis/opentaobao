package tmallnr

// NrtCrmBenefitDto 结构体
type NrtCrmBenefitDto struct {
	// 限领总额
	TotalQuantity string `json:"total_quantity,omitempty" xml:"total_quantity,omitempty"`
	// 个人限领
	PersonLimitCount string `json:"person_limit_count,omitempty" xml:"person_limit_count,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 门槛
	StartFee string `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 面额
	FaceValue string `json:"face_value,omitempty" xml:"face_value,omitempty"`
	// 券使用有效期间
	EffectiveDuration string `json:"effective_duration,omitempty" xml:"effective_duration,omitempty"`
	// 券使用结束时间
	EffectiveEndTime string `json:"effective_end_time,omitempty" xml:"effective_end_time,omitempty"`
	// 券使用开始时间
	EffectiveStartTime string `json:"effective_start_time,omitempty" xml:"effective_start_time,omitempty"`
	// 券使用时间类型
	EffectiveTimeMode string `json:"effective_time_mode,omitempty" xml:"effective_time_mode,omitempty"`
	// 券发送截止时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 券发送开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 券名称
	BenefitName string `json:"benefit_name,omitempty" xml:"benefit_name,omitempty"`
	// 券模板权益描述
	BenefitDescription string `json:"benefit_description,omitempty" xml:"benefit_description,omitempty"`
	// 券Code
	BenefitCode string `json:"benefit_code,omitempty" xml:"benefit_code,omitempty"`
	// 券模板实例Code
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
	// 券发放错误Code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 券发放错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 券模板实例ID
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 券模板ID
	OutId int64 `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 优惠券类型
	CouponType int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
}
