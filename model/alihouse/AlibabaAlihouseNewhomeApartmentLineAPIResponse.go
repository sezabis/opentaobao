package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeApartmentLineAPIResponse 公寓上下架 API返回值
// alibaba.alihouse.newhome.apartment.line
//
// 公寓上下架
type AlibabaAlihouseNewhomeApartmentLineAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeApartmentLineAPIResponseModel
}

// AlibabaAlihouseNewhomeApartmentLineAPIResponseModel is 公寓上下架 成功返回结果
type AlibabaAlihouseNewhomeApartmentLineAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_apartment_line_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeApartmentLineResult `json:"result,omitempty" xml:"result,omitempty"`
}
