package campus

// PojoResult 结构体
type PojoResult struct {
	// 内容
	Contents []CompanyDto `json:"contents,omitempty" xml:"contents>company_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 接口错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误级别
	ErrorLevel string `json:"error_level,omitempty" xml:"error_level,omitempty"`
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误详细信息
	ErrorExtInfo string `json:"error_ext_info,omitempty" xml:"error_ext_info,omitempty"`
	// 接口返回结果
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
