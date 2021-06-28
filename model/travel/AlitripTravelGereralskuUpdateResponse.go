package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发布SKU信息（如果properties重复 则更新） APIResponse
alitrip.travel.gereralsku.update

发布SKU信息（如果properties重复 则更新）
*/
type AlitripTravelGereralskuUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_travel_gereralsku_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    FirstResult   *TopTravelItem `json:"first_result,omitempty" xml:"