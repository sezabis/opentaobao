package cloudgame

// TopAvatarBodyDto 结构体
type TopAvatarBodyDto struct {
	// 头像icons
	AvatarIcons []string `json:"avatar_icons,omitempty" xml:"avatar_icons>string,omitempty"`
	// 装扮traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 装扮数据
	FaceDataJson string `json:"face_data_json,omitempty" xml:"face_data_json,omitempty"`
	// top请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 云游戏mixUserId
	MixUserId string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
	// 扩展数据
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 形象性别 1-男, 2-女
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
}
