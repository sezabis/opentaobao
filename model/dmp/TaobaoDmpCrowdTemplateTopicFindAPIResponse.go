package dmp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDmpCrowdTemplateTopicFindAPIResponse 平台精选榜单和模版查询接口 API返回值
// taobao.dmp.crowd.template.topic.find
//
// 查询平台精选榜单和模版信息
type TaobaoDmpCrowdTemplateTopicFindAPIResponse struct {
	model.CommonResponse
	TaobaoDmpCrowdTemplateTopicFindAPIResponseModel
}

// TaobaoDmpCrowdTemplateTopicFindAPIResponseModel is 平台精选榜单和模版查询接口 成功返回结果
type TaobaoDmpCrowdTemplateTopicFindAPIResponseModel struct {
	XMLName xml.Name `xml:"dmp_crowd_template_topic_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 榜单对象数组
	Result []Topic `json:"result,omitempty" xml:"result>topic,omitempty"`
	// 错误码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
