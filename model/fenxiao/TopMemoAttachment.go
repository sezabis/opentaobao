package fenxiao

// TopMemoAttachment 结构体
type TopMemoAttachment struct {
	// 附件地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 附件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
