package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
cp清理离职用户信息 APIResponse
cainiao.member.courier.cpresign

CP清理内部离职的用户信息
*/
type CainiaoMemberCourierCpresignAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_member_courier_cpresign_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 具体错误信息
    
    StatusMessage   string `json:"status_message,omitempty" xml:"