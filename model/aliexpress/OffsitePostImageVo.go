package aliexpress

// OffsitePostImageVo 结构体
type OffsitePostImageVo struct {
	// 图片链接
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 图片高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 图片宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}
