package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradesSimpleSoldGetAPIResponse 查询卖家已卖出的交易简易数据 API返回值
// taobao.trades.simple.sold.get
//
// 搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）
// <br/>1. 返回的数据结果是以订单的创建时间倒序排列的。
// <br/>2. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.simple.get获取订单信息。
// <br/>注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，解决办法就是type加上订单类型就可正常返回了。
// <br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong>
type TaobaoTradesSimpleSoldGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradesSimpleSoldGetAPIResponseModel
}

// TaobaoTradesSimpleSoldGetAPIResponseModel is 查询卖家已卖出的交易简易数据 成功返回结果
type TaobaoTradesSimpleSoldGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trades_simple_sold_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
	Trades []Trade `json:"trades,omitempty" xml:"trades>trade,omitempty"`
	// 搜索到的交易信息总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
