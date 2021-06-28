package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单创建接口 APIResponse
taobao.qimen.stockout.create

ERP调用奇门接口，创建出库单信息
*/
type TaobaoQimenStockoutCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_stockout_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"