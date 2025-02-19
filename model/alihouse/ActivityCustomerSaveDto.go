package alihouse

// ActivityCustomerSaveDto 结构体
type ActivityCustomerSaveDto struct {
	// 客户信息集合
	CustomerInfos []ActivityCustomerInfoDto `json:"customer_infos,omitempty" xml:"customer_infos>activity_customer_info_dto,omitempty"`
	// 外部项目店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部销售活动id
	OuterSalesActivityId string `json:"outer_sales_activity_id,omitempty" xml:"outer_sales_activity_id,omitempty"`
	// 外部楼盘id
	OuterProjectId string `json:"outer_project_id,omitempty" xml:"outer_project_id,omitempty"`
	// 请求时间版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
