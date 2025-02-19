package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServicedefinitionQueryskuAPIResponse 服务sku查询 API返回值
// alibaba.ssc.supplyplatform.servicedefinition.querysku
//
// 服务sku查询
type AlibabaSscSupplyplatformServicedefinitionQueryskuAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServicedefinitionQueryskuAPIResponseModel
}

// AlibabaSscSupplyplatformServicedefinitionQueryskuAPIResponseModel is 服务sku查询 成功返回结果
type AlibabaSscSupplyplatformServicedefinitionQueryskuAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_servicedefinition_querysku_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回类型
	Result *AlibabaSscSupplyplatformServicedefinitionQueryskuResult `json:"result,omitempty" xml:"result,omitempty"`
}
