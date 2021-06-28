package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取店铺关注链接 APIResponse
tmall.device.store.followurl.get

获取智能硬件上的关注店铺的URL
*/
type TmallDeviceStoreFollowurlGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_device_store_followurl_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 二维码短链接地址
    
    ShortUrl   string `json:"short_url,omitempty" xml:"