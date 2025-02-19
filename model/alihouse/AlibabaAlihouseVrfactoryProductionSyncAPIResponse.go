package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseVrfactoryProductionSyncAPIResponse vr生产数据上翻 API返回值
// alibaba.alihouse.vrfactory.production.sync
//
// vr生产数据上翻
type AlibabaAlihouseVrfactoryProductionSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseVrfactoryProductionSyncAPIResponseModel
}

// AlibabaAlihouseVrfactoryProductionSyncAPIResponseModel is vr生产数据上翻 成功返回结果
type AlibabaAlihouseVrfactoryProductionSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_vrfactory_production_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseVrfactoryProductionSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
