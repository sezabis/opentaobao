package ascp

// BatchCreateItemMappingResponse 结构体
type BatchCreateItemMappingResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务处理结果
	Data *BatchCreateItemMappingResult `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
