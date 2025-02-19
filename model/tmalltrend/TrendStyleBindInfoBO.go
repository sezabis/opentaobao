package tmalltrend

// TrendStyleBindInfoBO 结构体
type TrendStyleBindInfoBO struct {
	// 款式关联趋势词信息列表，趋势词信息：趋势词id+叶子类目id，单个款式最多关联3个趋势词
	TrendWordInfoList []string `json:"trend_word_info_list,omitempty" xml:"trend_word_info_list>string,omitempty"`
	// 同步操作目的，枚举，INSERT(&#34;新增&#34;),     UPDATE(&#34;更新&#34;),     OFFLINE(&#34;下线&#34;);
	SyncPurpose string `json:"sync_purpose,omitempty" xml:"sync_purpose,omitempty"`
	// 款式编号，业务唯一
	StyleSerialNumber string `json:"style_serial_number,omitempty" xml:"style_serial_number,omitempty"`
}
