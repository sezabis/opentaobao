package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证码列表查询 APIResponse
taobao.vmarket.eticket.codes.get

查询某个订单的所有码的列表
*/
type TaobaoVmarketEticketCodesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_codes_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 记录总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"