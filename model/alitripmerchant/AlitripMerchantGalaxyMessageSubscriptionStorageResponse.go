package alitripmerchant

// AlitripMerchantGalaxyMessageSubscriptionStorageResponse 结构体
type AlitripMerchantGalaxyMessageSubscriptionStorageResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 信息
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
