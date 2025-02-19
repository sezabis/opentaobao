package axindata

// MatchedRoomResultDto 结构体
type MatchedRoomResultDto struct {
	// 当前房型匹配最终结果列表，
	MatchedRoomDataList []MatchedRoomDataDto `json:"matched_room_data_list,omitempty" xml:"matched_room_data_list>matched_room_data_dto,omitempty"`
	// 房型名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 错误code
	ErCode string `json:"er_code,omitempty" xml:"er_code,omitempty"`
	// 错误信息
	ErMsg string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 当前房型的匹配是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
