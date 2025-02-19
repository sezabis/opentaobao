package ascpchannel

// PageData 结构体
type PageData struct {
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 商品名称
	ScItemName string `json:"sc_item_name,omitempty" xml:"sc_item_name,omitempty"`
	// 外部商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 创建时间
	CreateDate string `json:"create_date,omitempty" xml:"create_date,omitempty"`
	// 商品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}
