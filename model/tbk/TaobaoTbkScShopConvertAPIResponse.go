package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScShopConvertAPIResponse 淘宝客-服务商-店铺链接转换 API返回值
// taobao.tbk.sc.shop.convert
//
// 服务商使用。支持入参推广者对应的“推广位”和卖家id，获取对应的店铺推广链接。
type TaobaoTbkScShopConvertAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScShopConvertAPIResponseModel
}

// TaobaoTbkScShopConvertAPIResponseModel is 淘宝客-服务商-店铺链接转换 成功返回结果
type TaobaoTbkScShopConvertAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_shop_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 加入淘宝客的店铺
	Results []NTbkShop `json:"results,omitempty" xml:"results>n_tbk_shop,omitempty"`
}
