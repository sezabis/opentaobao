package btrip

// BaggageTip 结构体
type BaggageTip struct {
	// 行李子内容可视化内容
	BaggageSubContentVisualizes []BaggageSubContentVisualizesBean `json:"baggage_sub_content_visualizes,omitempty" xml:"baggage_sub_content_visualizes>baggage_sub_content_visualizes_bean,omitempty"`
	// 头像链接
	Logo string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 提示描述
	TipsDesc string `json:"tips_desc,omitempty" xml:"tips_desc,omitempty"`
	// 图片链接
	TipsImage string `json:"tips_image,omitempty" xml:"tips_image,omitempty"`
	// PTC
	Ptc string `json:"ptc,omitempty" xml:"ptc,omitempty"`
	// 退改签规则的类型
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 是否结构体
	IsStruct bool `json:"is_struct,omitempty" xml:"is_struct,omitempty"`
}
