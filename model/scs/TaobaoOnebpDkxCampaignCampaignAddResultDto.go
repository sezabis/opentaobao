package scs

// TaobaoOnebpDkxCampaignCampaignAddResultDto 结构体
type TaobaoOnebpDkxCampaignCampaignAddResultDto struct {
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	SolutionResultTopDTO *SolutionResultTopDto `json:"solution_result_top_d_t_o,omitempty" xml:"solution_result_top_d_t_o,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回状态码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
