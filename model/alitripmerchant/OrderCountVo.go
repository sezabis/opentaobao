package alitripmerchant

// OrderCountVo 结构体
type OrderCountVo struct {
	// 对应状态订单对应总数
	OrderCountDetailVOs []OrderCountDetailVo `json:"order_count_detail_v_os,omitempty" xml:"order_count_detail_v_os>order_count_detail_vo,omitempty"`
	// 酒店图片
	HotelPicture string `json:"hotel_picture,omitempty" xml:"hotel_picture,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 摘要
	AbstractContent string `json:"abstract_content,omitempty" xml:"abstract_content,omitempty"`
	// 入住日期
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 倒计时开始时间
	PayRemainTime int64 `json:"pay_remain_time,omitempty" xml:"pay_remain_time,omitempty"`
}
