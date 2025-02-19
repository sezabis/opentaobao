package axintrade

// RatePlanInfoApiDto 结构体
type RatePlanInfoApiDto struct {
	// 每间房rate信息
	RateUnitList []RateUnitDto `json:"rate_unit_list,omitempty" xml:"rate_unit_list>rate_unit_dto,omitempty"`
	// 最晚到店时间
	LatestCheckOutTime string `json:"latest_check_out_time,omitempty" xml:"latest_check_out_time,omitempty"`
	// 总房价
	TotalRoomPrice string `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
	// 床型描述
	BedTypeDesc string `json:"bed_type_desc,omitempty" xml:"bed_type_desc,omitempty"`
	// 最早可以办理入住时间
	EarliestCheckInTime string `json:"earliest_check_in_time,omitempty" xml:"earliest_check_in_time,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 最大订购数量
	MaxBookingNum int64 `json:"max_booking_num,omitempty" xml:"max_booking_num,omitempty"`
	// 每间房最大可入住人数
	MaxOccupancyNum int64 `json:"max_occupancy_num,omitempty" xml:"max_occupancy_num,omitempty"`
	// 取消政策
	CancelPolicy *CancelPolicyDto `json:"cancel_policy,omitempty" xml:"cancel_policy,omitempty"`
	// 最大库存量
	MaxInventory int64 `json:"max_inventory,omitempty" xml:"max_inventory,omitempty"`
	// 人民币总金额
	CnyTotalPrice int64 `json:"cny_total_price,omitempty" xml:"cny_total_price,omitempty"`
	// 汇率
	ExchangeRate *BigDecimal `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
	// 是否即时确认
	InstantConfirm bool `json:"instant_confirm,omitempty" xml:"instant_confirm,omitempty"`
}
