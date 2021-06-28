package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品编辑提交schema信息 APIResponse
alibaba.item.edit.submit

商品编辑提交schema信息
*/
type AlibabaItemEditSubmitAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_item_edit_submit_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品更新时间
    
    UpdateTime   string `json:"update_time,omitempty" xml:"