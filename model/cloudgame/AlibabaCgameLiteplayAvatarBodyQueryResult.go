package cloudgame

// AlibabaCgameLiteplayAvatarBodyQueryResult 结构体
type AlibabaCgameLiteplayAvatarBodyQueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据体
	Data *TopAvatarBodyDto `json:"data,omitempty" xml:"data,omitempty"`
}
