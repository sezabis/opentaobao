package tmallnr

// NrtCertificateInstanceQueryDto 结构体
type NrtCertificateInstanceQueryDto struct {
	// 淘系加密ID
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 电子凭证码值集合
	Codes string `json:"codes,omitempty" xml:"codes,omitempty"`
	// 业务编码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 显示数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
