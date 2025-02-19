package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbroomtypeDeleteAPIResponse 民宿房源删除接口 API返回值
// taobao.xhotel.bnbroomtype.delete
//
// 增加民宿房源删除接口
type TaobaoXhotelBnbroomtypeDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbroomtypeDeleteAPIResponseModel
}

// TaobaoXhotelBnbroomtypeDeleteAPIResponseModel is 民宿房源删除接口 成功返回结果
type TaobaoXhotelBnbroomtypeDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbroomtype_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否出错
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}
