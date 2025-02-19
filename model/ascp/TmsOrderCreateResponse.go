package ascp

// TmsOrderCreateResponse 结构体
type TmsOrderCreateResponse struct {
	// 业务错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 业务响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否为业务异常在流程中心场景下使用
	BizException bool `json:"biz_exception,omitempty" xml:"biz_exception,omitempty"`
	// 业务调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
