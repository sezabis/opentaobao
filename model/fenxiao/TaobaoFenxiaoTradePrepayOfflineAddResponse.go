package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下预存款流水增加 APIResponse
taobao.fenxiao.trade.prepay.offline.add

渠道分销供应商上传线下流水预存款（增加）
*/
type TaobaoFenxiaoTradePrepayOfflineAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_trade_prepay_offline_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultTopDo `json:"result,omitempty" xml:"