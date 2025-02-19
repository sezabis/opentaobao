package aecreatives

// TrafficFeaturedPromoResponseDto 结构体
type TrafficFeaturedPromoResponseDto struct {
	// 返回结果状态描述
	RespMsg string `json:"resp_msg,omitempty" xml:"resp_msg,omitempty"`
	// 返回结果状态码
	RespCode int64 `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 返回结果明细
	Result *TrafficFeaturedPromoResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
