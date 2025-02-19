package wdk

// AlibabaPricePromotionItemAddResult 结构体
type AlibabaPricePromotionItemAddResult struct {
	// 信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 编码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 结果内容
	Data *PromotionContentResult `json:"data,omitempty" xml:"data,omitempty"`
	// 数量
	TotalRecord int64 `json:"total_record,omitempty" xml:"total_record,omitempty"`
	// 是否成功
	SuccAndNotNull bool `json:"succ_and_not_null,omitempty" xml:"succ_and_not_null,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
