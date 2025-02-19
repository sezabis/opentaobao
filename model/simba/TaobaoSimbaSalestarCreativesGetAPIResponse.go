package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarCreativesGetAPIResponse （新）批量获取创意 API返回值
// taobao.simba.salestar.creatives.get
//
// 取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；&lt;br/&gt;如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
type TaobaoSimbaSalestarCreativesGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSalestarCreativesGetAPIResponseModel
}

// TaobaoSimbaSalestarCreativesGetAPIResponseModel is （新）批量获取创意 成功返回结果
type TaobaoSimbaSalestarCreativesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_creatives_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创意对象列表
	Creatives []Creative `json:"creatives,omitempty" xml:"creatives>creative,omitempty"`
}
