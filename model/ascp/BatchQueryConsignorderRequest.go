package ascp

// BatchQueryConsignorderRequest 结构体
type BatchQueryConsignorderRequest struct {
	// 业务请求ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间(毫秒数)
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 发货单最后修改时间区间
	LastModifiedPeriod *Period `json:"last_modified_period,omitempty" xml:"last_modified_period,omitempty"`
	// 每页数量，不大于100
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
}
