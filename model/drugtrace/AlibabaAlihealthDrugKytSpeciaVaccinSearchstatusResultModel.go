package drugtrace

// AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusResultModel 结构体
type AlibabaAlihealthDrugKytSpeciaVaccinSearchstatusResultModel struct {
	// 状态码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 响应结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
