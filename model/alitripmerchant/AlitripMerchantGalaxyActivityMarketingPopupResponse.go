package alitripmerchant

// AlitripMerchantGalaxyActivityMarketingPopupResponse 结构体
type AlitripMerchantGalaxyActivityMarketingPopupResponse struct {
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回结果
	Content *PopUpInfoVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
