package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取授权账号对应的OpenUid APIResponse
taobao.openuid.get

获取授权账号对应的OpenUid
*/
type TaobaoOpenuidGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openuid_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // OpenUID
    
    OpenUid   string `json:"open_uid,omitempty" xml:"