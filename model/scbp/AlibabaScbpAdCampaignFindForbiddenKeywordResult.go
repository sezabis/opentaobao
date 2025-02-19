package scbp

// AlibabaScbpAdCampaignFindForbiddenKeywordResult 结构体
type AlibabaScbpAdCampaignFindForbiddenKeywordResult struct {
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 普通词
	Normword string `json:"normword,omitempty" xml:"normword,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}
