package feedflow

// CampaignDto 结构体
type CampaignDto struct {
	// 打折范围
	LaunchPeriodList []LaunchPeriodDto `json:"launch_period_list,omitempty" xml:"launch_period_list>launch_period_dto,omitempty"`
	// 推广地域
	LaunchAreaList []LaunchAreaDto `json:"launch_area_list,omitempty" xml:"launch_area_list>launch_area_dto,omitempty"`
	// 计划名称
	CampaignName string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
	// PAUSE(&#34;投放暂停&#34;),START(&#34;投放开始&#34;),ERMINATE(&#34;投放停止&#34;),ABNORMAL(投放异常&#34;),WAIT(&#34;投放等待中&#34;),DELETE(&#34;删除&#34;)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 投放时间
	LaunchTime *LaunchTimeDto `json:"launch_time,omitempty" xml:"launch_time,omitempty"`
	// 每日预算，单位为分
	DayBudget int64 `json:"day_budget,omitempty" xml:"day_budget,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}
