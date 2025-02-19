package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGameGamepadGetAPIResponse 获取虚拟手柄配置 API返回值
// alibaba.cloudgame.interactive.game.gamepad.get
//
// 获取虚拟手柄配置
type AlibabaCloudgameInteractiveGameGamepadGetAPIResponse struct {
	model.CommonResponse
	AlibabaCloudgameInteractiveGameGamepadGetAPIResponseModel
}

// AlibabaCloudgameInteractiveGameGamepadGetAPIResponseModel is 获取虚拟手柄配置 成功返回结果
type AlibabaCloudgameInteractiveGameGamepadGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_gamepad_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *AlibabaCloudgameInteractiveGameGamepadGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
