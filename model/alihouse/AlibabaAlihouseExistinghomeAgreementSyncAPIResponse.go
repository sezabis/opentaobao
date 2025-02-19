package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeAgreementSyncAPIResponse 二手房电子协议数据同步 API返回值
// alibaba.alihouse.existinghome.agreement.sync
//
// 二手房电子协议数据同步
type AlibabaAlihouseExistinghomeAgreementSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeAgreementSyncAPIResponseModel
}

// AlibabaAlihouseExistinghomeAgreementSyncAPIResponseModel is 二手房电子协议数据同步 成功返回结果
type AlibabaAlihouseExistinghomeAgreementSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_agreement_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeAgreementSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
