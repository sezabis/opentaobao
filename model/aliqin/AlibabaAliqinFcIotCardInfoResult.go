package aliqin

// AlibabaAliqinFcIotCardInfoResult 结构体
type AlibabaAliqinFcIotCardInfoResult struct {
	// &#34;OpenTime&#34;:&#34;开户时间&#34;,&#34;IMSI&#34;:&#34;IMSI号&#34;,&#34;FirstActiveTime&#34;:&#34;第一次激活时间&#34;,&#34;MSISDN&#34;:&#34;MSISDN号&#34;,&#34;ICCID&#34;:&#34;ICCID号&#34;
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
