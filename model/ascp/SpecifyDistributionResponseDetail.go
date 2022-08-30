package ascp

// SpecifyDistributionResponseDetail 结构体
type SpecifyDistributionResponseDetail struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误消息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 要选择进行铺货的店铺宝贝 itemId
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 要选择进行铺货的店铺宝贝 skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 分销商sellerId
	DistributorShopUserId int64 `json:"distributor_shop_user_id,omitempty" xml:"distributor_shop_user_id,omitempty"`
	//  处理是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
