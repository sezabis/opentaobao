package ascp

// BatchUploadChannelInventoryRequest 结构体
type BatchUploadChannelInventoryRequest struct {
	// 渠道库存量，批量不超过50
	ChannelsInventory []ChannelInventory `json:"channels_inventory,omitempty" xml:"channels_inventory>channel_inventory,omitempty"`
	// 请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作人，用于记录操作记录
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 库存变动时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
