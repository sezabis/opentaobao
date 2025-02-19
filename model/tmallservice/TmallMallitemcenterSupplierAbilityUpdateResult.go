package tmallservice

// TmallMallitemcenterSupplierAbilityUpdateResult 结构体
type TmallMallitemcenterSupplierAbilityUpdateResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误类型
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 结果
	ResultData bool `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// true或false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
