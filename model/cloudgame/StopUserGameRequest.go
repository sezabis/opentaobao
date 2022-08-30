package cloudgame

// StopUserGameRequest 结构体
type StopUserGameRequest struct {
	// 游戏id
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}
