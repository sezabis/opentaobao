package pur

// AccessProductDto 结构体
type AccessProductDto struct {
	// 图片地址列表
	ImgUrlList []string `json:"img_url_list,omitempty" xml:"img_url_list>string,omitempty"`
	// 产品属性
	ProductAttrValueList []AccessProductAttrValueDto `json:"product_attr_value_list,omitempty" xml:"product_attr_value_list>access_product_attr_value_dto,omitempty"`
	// 阿里采购员工号
	BuyerWorkNo string `json:"buyer_work_no,omitempty" xml:"buyer_work_no,omitempty"`
	// 阿里采购三级类目
	CategoryCode string `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// 外部商家标识
	DataSource string `json:"data_source,omitempty" xml:"data_source,omitempty"`
	// 外部商家产品链接
	MallUrl string `json:"mall_url,omitempty" xml:"mall_url,omitempty"`
	// 计价方式 AMOUNT QUANTITY
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 产品描述
	ProductDesc string `json:"product_desc,omitempty" xml:"product_desc,omitempty"`
	// 产品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 采购通道，阿里侧枚举
	PurchaseChannel string `json:"purchase_channel,omitempty" xml:"purchase_channel,omitempty"`
	// 产品详情
	Recommendation string `json:"recommendation,omitempty" xml:"recommendation,omitempty"`
	// 外部商家品类名称
	SourceCategoryName string `json:"source_category_name,omitempty" xml:"source_category_name,omitempty"`
	// 额外信息
	SourceInfo string `json:"source_info,omitempty" xml:"source_info,omitempty"`
	// 子类型
	SourceType string `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// 外部商家对应产品ID
	SourceValue string `json:"source_value,omitempty" xml:"source_value,omitempty"`
	// 计价单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 图片存储方式
	ImgStorageType string `json:"img_storage_type,omitempty" xml:"img_storage_type,omitempty"`
	// 阿里商城品类ID
	CatalogCategoryId int64 `json:"catalog_category_id,omitempty" xml:"catalog_category_id,omitempty"`
	// 外部商家品类ID，如果有的话需要在阿里侧有对应的映射
	SourceCategoryId int64 `json:"source_category_id,omitempty" xml:"source_category_id,omitempty"`
	// 租户
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 采购类别
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 品牌ID
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 采购用途
	CategoryUseId int64 `json:"category_use_id,omitempty" xml:"category_use_id,omitempty"`
	// 前台类目
	FrontCategoryId int64 `json:"front_category_id,omitempty" xml:"front_category_id,omitempty"`
}
