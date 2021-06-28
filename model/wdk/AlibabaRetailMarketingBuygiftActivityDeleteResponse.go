package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单品买赠活动 APIResponse
alibaba.retail.marketing.buygift.activity.delete

同城零售单品特价活动删除
*/
type AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_retail_marketing_buygift_activity_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    Result   *OctopusOpenResult `json:"result,omitempty" xml:"