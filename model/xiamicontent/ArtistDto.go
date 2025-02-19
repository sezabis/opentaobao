package xiamicontent

// ArtistDto 结构体
type ArtistDto struct {
	// 地区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 性别
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
	// 别名
	Alias string `json:"alias,omitempty" xml:"alias,omitempty"`
	// 艺人名
	ArtistName string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
	// 艺人图片
	ArtistLogo string `json:"artist_logo,omitempty" xml:"artist_logo,omitempty"`
	// 艺人id
	ArtistId int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
}
