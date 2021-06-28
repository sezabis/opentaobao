package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客户端授权页 APIResponse
alibaba.interact.sensor.authorize

客户端授权页
*/
type AlibabaInteractSensorAuthorizeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_authorize_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // return=0 表示成功
    
    Result   string `json:"result,omitempty" xml:"