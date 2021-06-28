package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定设备 APIResponse
alibaba.alink.device.bind

阿里智能解绑设备
*/
type AlibabaAlinkDeviceBindAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alink_device_bind_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *TopServiceResult `json:"result,omitempty" xml:"