package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseImReceiveModelSyncAPIResponse IM承接方式同步 API返回值
// alibaba.alihouse.im.receive.model.sync
//
// IM承接方式同步
type AlibabaAlihouseImReceiveModelSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseImReceiveModelSyncAPIResponseModel
}

// AlibabaAlihouseImReceiveModelSyncAPIResponseModel is IM承接方式同步 成功返回结果
type AlibabaAlihouseImReceiveModelSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_im_receive_model_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// aaa
	Result *AlibabaAlihouseImReceiveModelSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
