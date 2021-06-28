package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询储值流水 APIResponse
alibaba.alsc.crm.recharge.accountflows.get

增加分页查询储值流水接口
*/
type AlibabaAlscCrmRechargeAccountflowsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_recharge_accountflows_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分页返回模型
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"