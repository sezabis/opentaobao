package axintrade

// BoardDto 结构体
type BoardDto struct {
	// 餐食数量
	BoardNum int64 `json:"board_num,omitempty" xml:"board_num,omitempty"`
	// 餐食种类
	BoardType int64 `json:"board_type,omitempty" xml:"board_type,omitempty"`
}
