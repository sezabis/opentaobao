package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappWidgetTemplateInstanceQueryAPIResponse 小部件实例化版本查询 API返回值
// taobao.miniapp.widget.template.instance.query
//
// 小部件实例化版本查询
type TaobaoMiniappWidgetTemplateInstanceQueryAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappWidgetTemplateInstanceQueryAPIResponseModel
}

// TaobaoMiniappWidgetTemplateInstanceQueryAPIResponseModel is 小部件实例化版本查询 成功返回结果
type TaobaoMiniappWidgetTemplateInstanceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_widget_template_instance_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoMiniappWidgetTemplateInstanceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
