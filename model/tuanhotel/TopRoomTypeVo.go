package tuanhotel

// TopRoomTypeVo 结构体
type TopRoomTypeVo struct {
	// 房型名称
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// 房型ID
	TypeId int64 `json:"type_id,omitempty" xml:"type_id,omitempty"`
}
