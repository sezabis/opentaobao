package mos

// SmsSendMessageDto 结构体
type SmsSendMessageDto struct {
	// 手机号
	PhoneNumber string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 租户
	MosTenant string `json:"mos_tenant,omitempty" xml:"mos_tenant,omitempty"`
	// 短信模板ID
	SmsTemplateId string `json:"sms_template_id,omitempty" xml:"sms_template_id,omitempty"`
	// 幂等ID防重试
	RelationId string `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 当前租户短信使用场景
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 模板参数填充MAP
	TemplateParamsMap string `json:"template_params_map,omitempty" xml:"template_params_map,omitempty"`
}
