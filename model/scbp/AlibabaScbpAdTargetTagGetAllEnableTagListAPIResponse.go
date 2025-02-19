package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagGetAllEnableTagListAPIResponse 查询所有可添加标签信息 API返回值
// alibaba.scbp.ad.target.tag.get.all.enable.tag.list
//
// 查询标签数据
type AlibabaScbpAdTargetTagGetAllEnableTagListAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdTargetTagGetAllEnableTagListAPIResponseModel
}

// AlibabaScbpAdTargetTagGetAllEnableTagListAPIResponseModel is 查询所有可添加标签信息 成功返回结果
type AlibabaScbpAdTargetTagGetAllEnableTagListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_target_tag_get_all_enable_tag_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回实体
	ResultList []AdsTargetingTagDto `json:"result_list,omitempty" xml:"result_list>ads_targeting_tag_dto,omitempty"`
}
