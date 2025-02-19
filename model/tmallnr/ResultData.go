package tmallnr

// ResultData 结构体
type ResultData struct {
	// 服务范围，支持多服务范围返回
	Ranges []NrServiceRangeResponseDto `json:"ranges,omitempty" xml:"ranges>nr_service_range_response_dto,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
