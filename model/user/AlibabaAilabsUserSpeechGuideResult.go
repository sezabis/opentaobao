package user

// AlibabaAilabsUserSpeechGuideResult 结构体
type AlibabaAilabsUserSpeechGuideResult struct {
	// 出错信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 0表示成功
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 推荐信息model
	RetValue *RecommendInfo `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
}
