package cainiaohandover

// HsfResult 结构体
type HsfResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 异常码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 响应数据
	Data *OpenHandoverCancelResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 请求处理是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 请求结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
