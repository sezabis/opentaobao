package cloudgame

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameInteractiveGamePlayerKickoutAPIResponse 踢人 API返回值
// alibaba.cloudgame.interactive.game.player.kickout
//
// 踢人
type AlibabaCloudgameInteractiveGamePlayerKickoutAPIResponse struct {
	model.CommonResponse
	AlibabaCloudgameInteractiveGamePlayerKickoutAPIResponseModel
}

// AlibabaCloudgameInteractiveGamePlayerKickoutAPIResponseModel is 踢人 成功返回结果
type AlibabaCloudgameInteractiveGamePlayerKickoutAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cloudgame_interactive_game_player_kickout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaCloudgameInteractiveGamePlayerKickoutResult `json:"result,omitempty" xml:"result,omitempty"`
}
