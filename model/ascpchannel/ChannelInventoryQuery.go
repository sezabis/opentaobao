package ascpchannel

// ChannelInventoryQuery 结构体
type ChannelInventoryQuery struct {
	// 产品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 二级渠道
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 默认不需传
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 一级渠道
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 省
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
}
