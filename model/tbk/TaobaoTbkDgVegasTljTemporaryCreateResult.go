package tbk

// TaobaoTbkDgVegasTljTemporaryCreateResult 结构体
type TaobaoTbkDgVegasTljTemporaryCreateResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *RightsInstanceCreateResult `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
