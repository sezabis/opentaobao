package fenxiao

// CnskuResult 结构体
type CnskuResult struct {
	// 对应货品信息
	Data []CnskuDto `json:"data,omitempty" xml:"data>cnsku_dto,omitempty"`
	// 异常信息Code
	SysErrorCode string `json:"sys_error_code,omitempty" xml:"sys_error_code,omitempty"`
	// 异常信息
	ErrorMSG string `json:"error_m_s_g,omitempty" xml:"error_m_s_g,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否系统异常
	IsSystemFailed bool `json:"is_system_failed,omitempty" xml:"is_system_failed,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
