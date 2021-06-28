package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除通用单品优惠详情 APIResponse
taobao.promotionmisc.common.item.detail.delete

删除通用单品优惠详情。
*/
type TaobaoPromotionmiscCommonItemDetailDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotionmisc_common_item_detail_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否删除成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"