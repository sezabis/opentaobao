package tbk

// TaobaoTbkScRelationRefundRpcResult 结构体
type TaobaoTbkScRelationRefundRpcResult struct {
	// 返回信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 业务错误信息
	BizErrorDesc string `json:"biz_error_desc,omitempty" xml:"biz_error_desc,omitempty"`
	// 真正的业务数据结构
	Data *PageResult `json:"data,omitempty" xml:"data,omitempty"`
	// 接口返回值信息，跟rpc架构保持一致
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 业务错误码 101, 102,103
	BizErrorCode int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
}
