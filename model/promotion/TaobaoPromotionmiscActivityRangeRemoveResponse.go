package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
去除活动参与的商品 APIResponse
taobao.promotionmisc.activity.range.remove

去除活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeRemoveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotionmisc_activity_range_remove_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 去除活动参与的商品是否成功。
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"