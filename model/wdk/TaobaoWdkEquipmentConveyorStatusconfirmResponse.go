package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
悬挂链状态回传确认 APIResponse
taobao.wdk.equipment.conveyor.statusconfirm

悬挂链状态回传确认
*/
type TaobaoWdkEquipmentConveyorStatusconfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wdk_equipment_conveyor_statusconfirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TopWcsResult `json:"result,omitempty" xml:"