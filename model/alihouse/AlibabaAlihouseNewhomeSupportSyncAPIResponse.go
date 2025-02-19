package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeSupportSyncAPIResponse 周边配套数据同步 API返回值
// alibaba.alihouse.newhome.support.sync
//
// 周边配套数据同步
type AlibabaAlihouseNewhomeSupportSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeSupportSyncAPIResponseModel
}

// AlibabaAlihouseNewhomeSupportSyncAPIResponseModel is 周边配套数据同步 成功返回结果
type AlibabaAlihouseNewhomeSupportSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_support_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeSupportSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
