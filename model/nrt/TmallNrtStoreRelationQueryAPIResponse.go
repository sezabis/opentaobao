package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStoreRelationQueryAPIResponse 喵零门店关系查询 API返回值
// tmall.nrt.store.relation.query
//
// 喵零门店关系查询
type TmallNrtStoreRelationQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtStoreRelationQueryAPIResponseModel
}

// TmallNrtStoreRelationQueryAPIResponseModel is 喵零门店关系查询 成功返回结果
type TmallNrtStoreRelationQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_store_relation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
