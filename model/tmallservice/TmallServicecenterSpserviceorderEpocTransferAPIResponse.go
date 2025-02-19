package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderEpocTransferAPIResponse 电子保单受保人转移 API返回值
// tmall.servicecenter.spserviceorder.epoc.transfer
//
// 电子保单受保人转移
type TmallServicecenterSpserviceorderEpocTransferAPIResponse struct {
	model.CommonResponse
	TmallServicecenterSpserviceorderEpocTransferAPIResponseModel
}

// TmallServicecenterSpserviceorderEpocTransferAPIResponseModel is 电子保单受保人转移 成功返回结果
type TmallServicecenterSpserviceorderEpocTransferAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_spserviceorder_epoc_transfer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
