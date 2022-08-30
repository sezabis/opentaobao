package scs

// TaobaoOnebpDkxCampaignCampaignNoreportResultDto 结构体
type TaobaoOnebpDkxCampaignCampaignNoreportResultDto struct {
	// 返回结果
	CampaignResultTopDTOList []CampaignResultTopDto `json:"campaign_result_top_d_t_o_list,omitempty" xml:"campaign_result_top_d_t_o_list>campaign_result_top_dto,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回状态码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
