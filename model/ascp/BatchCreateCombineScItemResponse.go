package ascp

// BatchCreateCombineScItemResponse 结构体
type BatchCreateCombineScItemResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务处理结果
	Data *BatchCreateCombineScItemResult `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
