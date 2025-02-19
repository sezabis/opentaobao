package btrip

// PassengerCabinInfo 结构体
type PassengerCabinInfo struct {
	// 原航班号
	OriginFlightNo string `json:"origin_flight_no,omitempty" xml:"origin_flight_no,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 用户编号
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 改签费
	ChangeFee int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// 升舱费用
	UpgradeFee int64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}
