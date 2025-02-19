package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolSkuQueryAPIResponse 查询商品池活动商品【同城零售】 API返回值
// alibaba.retail.marketing.itempool.sku.query
//
// 查询商品池活动商品【同城零售】
type AlibabaRetailMarketingItempoolSkuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItempoolSkuQueryAPIResponseModel
}

// AlibabaRetailMarketingItempoolSkuQueryAPIResponseModel is 查询商品池活动商品【同城零售】 成功返回结果
type AlibabaRetailMarketingItempoolSkuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_sku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应体
	Data []ItemPoolSkuActivityElementDto `json:"data,omitempty" xml:"data>item_pool_sku_activity_element_dto,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 错误码
	ErrNumber string `json:"err_number,omitempty" xml:"err_number,omitempty"`
	// 分页信息
	PageInfo *PageInfoDto `json:"page_info,omitempty" xml:"page_info,omitempty"`
	// 成功标识
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}
