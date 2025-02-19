package util

// AiotOpenDeviceBaseDto 结构体
type AiotOpenDeviceBaseDto struct {
	// 设备业务标识
	Utdid string `json:"utdid,omitempty" xml:"utdid,omitempty"`
	// 设备名
	DeviceName string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	// 厂商
	Manufacturer string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// 品牌
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 硬件型号
	HardCode string `json:"hard_code,omitempty" xml:"hard_code,omitempty"`
	// 设备类型
	DeviceType string `json:"device_type,omitempty" xml:"device_type,omitempty"`
	// 设备SN
	DeviceSn string `json:"device_sn,omitempty" xml:"device_sn,omitempty"`
	// 平台类型
	Platform int64 `json:"platform,omitempty" xml:"platform,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
