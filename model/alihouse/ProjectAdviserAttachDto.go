package alihouse

// ProjectAdviserAttachDto 结构体
type ProjectAdviserAttachDto struct {
	// 近90天带看量
	RecentViewCount int64 `json:"recent_view_count,omitempty" xml:"recent_view_count,omitempty"`
	// 考试认证1-已认证 0-未认证
	ExamStatus int64 `json:"exam_status,omitempty" xml:"exam_status,omitempty"`
	// 导师认证1-已认证 0-未认证
	TutorStatus int64 `json:"tutor_status,omitempty" xml:"tutor_status,omitempty"`
	// 近90天交易量
	RecentTradeCount int64 `json:"recent_trade_count,omitempty" xml:"recent_trade_count,omitempty"`
}
