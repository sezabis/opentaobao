package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionOrderItemsAllBindAPIResponse 小程序投放-基于投放计划绑定/解绑全店商品 API返回值
// taobao.miniapp.distribution.order.items.all.bind
//
// 提供给使用了投放插件的服务商，可以调用该API实现帮助商家更新已创建的投放单中的绑定商品信息。该接口对应的能力是全店投放，表示店铺里的所有商品均会透出对应的投放计划（并且对于新上架的商品，系统也会每天凌晨进行一次补充投放，始终保持全店商品都有悬浮窗入口），该能力的使用，需要联系平台运营进行人工申请，申请通过以后方可调用成功。
type TaobaoMiniappDistributionOrderItemsAllBindAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionOrderItemsAllBindAPIResponseModel
}

// TaobaoMiniappDistributionOrderItemsAllBindAPIResponseModel is 小程序投放-基于投放计划绑定/解绑全店商品 成功返回结果
type TaobaoMiniappDistributionOrderItemsAllBindAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_order_items_all_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否绑定/解绑成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
