package eticket

// EticketCode 结构体
type EticketCode struct {
	// 电子凭证码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 二维码的图片地址
	QrcodeUrl string `json:"qrcode_url,omitempty" xml:"qrcode_url,omitempty"`
	// 可用数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 码状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
