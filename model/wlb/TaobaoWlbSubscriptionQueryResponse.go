package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家定购的所有服务 APIResponse
taobao.wlb.subscription.query

查询商家定购的所有服务,可通过入参状态来筛选
*/
type TaobaoWlbSubscriptionQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_subscription_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果总数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"