package cloudgame

// JoinCodeAssignResponse 结构体
type JoinCodeAssignResponse struct {
	// 玩家列表
	PlayerList []Long `json:"player_list,omitempty" xml:"player_list>long,omitempty"`
	// 联机码
	JoinCode string `json:"join_code,omitempty" xml:"join_code,omitempty"`
	// 游戏会话id
	GameSession string `json:"game_session,omitempty" xml:"game_session,omitempty"`
	// 扩展字段
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 游戏详情
	Game *OpenGameDto `json:"game,omitempty" xml:"game,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}
