package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreSyncAPIResponse 二手房标准门店数据同步 API返回值
// alibaba.alihouse.existinghome.store.sync
//
// 二手房标准门店数据同步
type AlibabaAlihouseExistinghomeStoreSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeStoreSyncAPIResponseModel
}

// AlibabaAlihouseExistinghomeStoreSyncAPIResponseModel is 二手房标准门店数据同步 成功返回结果
type AlibabaAlihouseExistinghomeStoreSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_store_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeStoreSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
