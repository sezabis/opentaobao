package alitripmerchant

// AlitripMerchantGalaxyVerifySignatureResponse 结构体
type AlitripMerchantGalaxyVerifySignatureResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回内容
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}
