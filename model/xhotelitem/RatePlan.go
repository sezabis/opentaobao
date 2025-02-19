package xhotelitem

// RatePlan 结构体
type RatePlan struct {
	// 系统商，一般不填写，使用须申请
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 房价名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 卖家自己系统的Code，简称RateCode
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
	// 卖家自己系统的Code，简称RateCode
	RateplanCode string `json:"rateplan_code,omitempty" xml:"rateplan_code,omitempty"`
	// 英文名称
	EnglishName string `json:"english_name,omitempty" xml:"english_name,omitempty"`
	// 额外服务的扩展，是一段JSON
	ExtendFee string `json:"extend_fee,omitempty" xml:"extend_fee,omitempty"`
	// 每日生效时间默认00:00:00。生效时间＜结束时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 每日结束时间默认24:00:00。生效时间＜结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 退订政策
	CancelPolicy string `json:"cancel_policy,omitempty" xml:"cancel_policy,omitempty"`
	// extend
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 扩展字段1
	ExtendInfo1 string `json:"extend_info1,omitempty" xml:"extend_info1,omitempty"`
	// 扩展字段2
	ExtendInfo2 string `json:"extend_info2,omitempty" xml:"extend_info2,omitempty"`
	// 扩展字段3
	ExtendInfo3 string `json:"extend_info3,omitempty" xml:"extend_info3,omitempty"`
	// 每日开始担保时间
	GuaranteeStartTime string `json:"guarantee_start_time,omitempty" xml:"guarantee_start_time,omitempty"`
	// 会员等级。支持多个等级&#34;,&#34;分隔
	MemberLevel string `json:"member_level,omitempty" xml:"member_level,omitempty"`
	// 销售渠道，目前制定一了一种A-集团协议
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// rateplan生效开始时间
	EffectiveTime string `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
	// rateplan生效截止时间
	DeadlineTime string `json:"deadline_time,omitempty" xml:"deadline_time,omitempty"`
	// 协议保留房提前确认时间
	AllotmentReleaseTime string `json:"allotment_release_time,omitempty" xml:"allotment_release_time,omitempty"`
	// rp维度的发票信息,type:1.酒店提供发票；2.卖家提供发票,desc:发票描叙，比如：卖家包邮提供发票,格式为：{&#34;type&#34;:1;&#34;desc&#34;:&#34;&#34;}
	InvoiceContent string `json:"invoice_content,omitempty" xml:"invoice_content,omitempty"`
	// key的含义:    non-direct-RP 表示非直连RP,      super-could-price-change-RP 表示rp的super标，打上这个tag，表明这个rateplan下单的时候支持变价，商家系统直接放开价格校验。      base-could-derived-RP 表示base rateplan标签，打上了这个tag，表明这是一个base的rateplan，基于该rateplan可以衍生出子rateplan  .        ebk-tail-room-RP 表示 ebk尾房rate plan级别标
	TagJson string `json:"tag_json,omitempty" xml:"tag_json,omitempty"`
	// 可入住的最晚时间（小时房相关字段）
	CanCheckinEnd string `json:"can_checkin_end,omitempty" xml:"can_checkin_end,omitempty"`
	// 可入住的最早时间（小时房相关字段）
	CanCheckinStart string `json:"can_checkin_start,omitempty" xml:"can_checkin_start,omitempty"`
	// 每日生效结束时间（仅时分秒有效）
	EndTimeDaily string `json:"end_time_daily,omitempty" xml:"end_time_daily,omitempty"`
	// rateplan类型 1为小时房
	RpType string `json:"rp_type,omitempty" xml:"rp_type,omitempty"`
	// 入住的开始跨度（小时房相关字段）
	Hourage string `json:"hourage,omitempty" xml:"hourage,omitempty"`
	// 每日生效开始时间（仅时分秒有效）
	StartTimeDaily string `json:"start_time_daily,omitempty" xml:"start_time_daily,omitempty"`
	// 普通保留房提前确认时间
	CommonAllotReleaseTime string `json:"common_allot_release_time,omitempty" xml:"common_allot_release_time,omitempty"`
	// hotelCompanyMappingDOS
	HotelCompanyMappingDOS string `json:"hotel_company_mapping_d_o_s,omitempty" xml:"hotel_company_mapping_d_o_s,omitempty"`
	// calBreakfastStr
	CalBreakfastStr string `json:"cal_breakfast_str,omitempty" xml:"cal_breakfast_str,omitempty"`
	// calGuaranteeStr
	CalGuaranteeStr string `json:"cal_guarantee_str,omitempty" xml:"cal_guarantee_str,omitempty"`
	// calChangeRuleStr
	CalChangeRuleStr string `json:"cal_change_rule_str,omitempty" xml:"cal_change_rule_str,omitempty"`
	// 可离店的最晚时间（小时房相关字段）
	CanCheckoutEnd string `json:"can_checkout_end,omitempty" xml:"can_checkout_end,omitempty"`
	// 会员价加价规则。c:表示折扣百分比,例子8,意为会员价优惠8%,s:标识起始日期,e:表示截止日期，t:表示加价类型，0:代表折扣。会员价=售价*(1-c%)
	MemberDiscountCal string `json:"member_discount_cal,omitempty" xml:"member_discount_cal,omitempty"`
	// 酒+X特色
	Benefits string `json:"benefits,omitempty" xml:"benefits,omitempty"`
	// 活动类型: 1通兑,2秒杀,3尾房,4超级房券
	ActivityType string `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 阿里房价id
	RatePlanId int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	// 1：开启2：关闭。
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// rateplan_id
	Rpid int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`
	// 支付类型 可选值 1：预付 5：前台面付
	PaymentType int64 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// 早餐数量
	BreakfastCount int64 `json:"breakfast_count,omitempty" xml:"breakfast_count,omitempty"`
	// 另加早餐数量
	FeeBreakfastCount int64 `json:"fee_breakfast_count,omitempty" xml:"fee_breakfast_count,omitempty"`
	// 另加早餐金额
	FeeBreakfastAmount int64 `json:"fee_breakfast_amount,omitempty" xml:"fee_breakfast_amount,omitempty"`
	// 额外服务-政府税-金额（1-9999）
	FeeGovTaxAmount int64 `json:"fee_gov_tax_amount,omitempty" xml:"fee_gov_tax_amount,omitempty"`
	// 额外服务-政府税-百分比（0%-99%）
	FeeGovTaxPercent int64 `json:"fee_gov_tax_percent,omitempty" xml:"fee_gov_tax_percent,omitempty"`
	// 额外服务-服务费-金额（0-9999）
	FeeServiceAmount int64 `json:"fee_service_amount,omitempty" xml:"fee_service_amount,omitempty"`
	// 额外服务-服务费-百分比（0%-99%）
	FeeServicePercent int64 `json:"fee_service_percent,omitempty" xml:"fee_service_percent,omitempty"`
	// 最小入住天数（1-365）
	MinDays int64 `json:"min_days,omitempty" xml:"min_days,omitempty"`
	// 最大入住天数（1-365）
	MaxDays int64 `json:"max_days,omitempty" xml:"max_days,omitempty"`
	// 首日入住房间数（1-99）【废弃】
	MinAmount int64 `json:"min_amount,omitempty" xml:"min_amount,omitempty"`
	// 最小提前预订小时按入住时间的23:59:59(一般认为24点)来计算
	MinAdvHours int64 `json:"min_adv_hours,omitempty" xml:"min_adv_hours,omitempty"`
	// 最大提前预订小时按入住时间的23:59:59(一般认为24点)来计算
	MaxAdvHours int64 `json:"max_adv_hours,omitempty" xml:"max_adv_hours,omitempty"`
	// 担保类型，只支持： 0 无担保 1 峰时首晚担保 2峰时全额担保 3全天首晚担保 4全天全额担保
	GuaranteeType int64 `json:"guarantee_type,omitempty" xml:"guarantee_type,omitempty"`
	// 入住人数
	Occupancy int64 `json:"occupancy,omitempty" xml:"occupancy,omitempty"`
	// 是否是首住优惠rp。1代表是
	FirstStay int64 `json:"first_stay,omitempty" xml:"first_stay,omitempty"`
	// 是否是协议价。1代表是
	Agreement int64 `json:"agreement,omitempty" xml:"agreement,omitempty"`
	// 0支付宝担保 1PCI担保
	GuaranteeMode int64 `json:"guarantee_mode,omitempty" xml:"guarantee_mode,omitempty"`
	// 是否包房RP 1包房RP,0 非包房rp
	PackRoomFlag int64 `json:"pack_room_flag,omitempty" xml:"pack_room_flag,omitempty"`
	// 是否底价加价，1是底价加价,0 非底价加价rp
	BottomPriceFlag int64 `json:"bottom_price_flag,omitempty" xml:"bottom_price_flag,omitempty"`
	// 是否为学生价
	IsStudent int64 `json:"is_student,omitempty" xml:"is_student,omitempty"`
	// 来源
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// sell 端特殊RP 对应的 gid
	SellGid int64 `json:"sell_gid,omitempty" xml:"sell_gid,omitempty"`
	// 父rpid
	ParentRpid int64 `json:"parent_rpid,omitempty" xml:"parent_rpid,omitempty"`
	// companyAssist
	CompanyAssist int64 `json:"company_assist,omitempty" xml:"company_assist,omitempty"`
	// 会员价支持标识,1表示支持会员价规则
	MemDiscFlag int64 `json:"mem_disc_flag,omitempty" xml:"mem_disc_flag,omitempty"`
}
