package subuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubusersInfoQueryAPIResponse 根据当前子账号登陆态，获取该子账号基本信息 API返回值
// taobao.subusers.info.query
//
// 根据当前子账号登陆态，获取该子账号基本信息
type TaobaoSubusersInfoQueryAPIResponse struct {
	model.CommonResponse
	TaobaoSubusersInfoQueryAPIResponseModel
}

// TaobaoSubusersInfoQueryAPIResponseModel is 根据当前子账号登陆态，获取该子账号基本信息 成功返回结果
type TaobaoSubusersInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"subusers_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号对象，如果为空，说明没查到该子账号
	Result *SubUserDo `json:"result,omitempty" xml:"result,omitempty"`
}
