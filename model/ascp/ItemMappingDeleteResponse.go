package ascp

// ItemMappingDeleteResponse 结构体
type ItemMappingDeleteResponse struct {
	// 业务处理结果
	Data []DataItem `json:"data,omitempty" xml:"data>data_item,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
