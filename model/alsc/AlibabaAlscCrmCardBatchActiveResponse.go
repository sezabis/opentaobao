package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量激活卡 APIResponse
alibaba.alsc.crm.card.batch.active

批量激活卡
*/
type AlibabaAlscCrmCardBatchActiveAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCardBatchActiveResponse
}

type AlibabaAlscCrmCardBatchActiveResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_card_batch_active_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
