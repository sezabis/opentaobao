package category

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemcatsGetAPIResponse 获取后台供卖家发布商品的标准商品类目 API返回值
// taobao.itemcats.get
//
// 获取后台供卖家发布商品的标准商品类目。
type TaobaoItemcatsGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemcatsGetAPIResponseModel
}

// TaobaoItemcatsGetAPIResponseModel is 获取后台供卖家发布商品的标准商品类目 成功返回结果
type TaobaoItemcatsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"itemcats_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品类目结构
	ItemCats []ItemCat `json:"item_cats,omitempty" xml:"item_cats>item_cat,omitempty"`
	// 最近修改时间(如果取增量，会返回该字段)。
	LastModified string `json:"last_modified,omitempty" xml:"last_modified,omitempty"`
}
