package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgCpaActivityReportAPIResponse 淘宝客-推广者-任务奖励效果报表 API返回值
// taobao.tbk.dg.cpa.activity.report
//
// 提供给媒体使用的cpa活动报表查询api，当前仅试运行媒体可使用
type TaobaoTbkDgCpaActivityReportAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgCpaActivityReportAPIResponseModel
}

// TaobaoTbkDgCpaActivityReportAPIResponseModel is 淘宝客-推广者-任务奖励效果报表 成功返回结果
type TaobaoTbkDgCpaActivityReportAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_cpa_activity_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回模型
	Result *TaobaoTbkDgCpaActivityReportRpcResult `json:"result,omitempty" xml:"result,omitempty"`
}
