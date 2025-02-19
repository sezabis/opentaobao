package alihealth2

// TopResultModel 结构体
type TopResultModel struct {
	// 返回json字符串，包含商家的省、市、县、和id
	Models []string `json:"models,omitempty" xml:"models>string,omitempty"`
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// isvId
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}
