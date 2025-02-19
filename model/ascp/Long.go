package ascp

// Long 结构体
type Long struct {
	// 币种
	DistributeCurrency string `json:"distribute_currency,omitempty" xml:"distribute_currency,omitempty"`
	// 币种
	RetailCurrency string `json:"retail_currency,omitempty" xml:"retail_currency,omitempty"`
	// 【必传】分销商sellerId
	DistributorShopUserId int64 `json:"distributor_shop_user_id,omitempty" xml:"distributor_shop_user_id,omitempty"`
	// 【必传】分销价
	DistributePrice int64 `json:"distribute_price,omitempty" xml:"distribute_price,omitempty"`
	// 建议零售价
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
}
