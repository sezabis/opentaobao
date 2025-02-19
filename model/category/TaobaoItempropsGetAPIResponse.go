package category

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItempropsGetAPIResponse 获取标准商品类目属性 API返回值
// taobao.itemprops.get
//
// 通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。
type TaobaoItempropsGetAPIResponse struct {
	model.CommonResponse
	TaobaoItempropsGetAPIResponseModel
}

// TaobaoItempropsGetAPIResponseModel is 获取标准商品类目属性 成功返回结果
type TaobaoItempropsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"itemprops_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品属性
	ItemProps []ItemProp `json:"item_props,omitempty" xml:"item_props>item_prop,omitempty"`
	// 最近修改时间(只有取全量或增量的时候会返回该字段)。格式:yyyy-MM-dd HH:mm:ss
	LastModified string `json:"last_modified,omitempty" xml:"last_modified,omitempty"`
}
