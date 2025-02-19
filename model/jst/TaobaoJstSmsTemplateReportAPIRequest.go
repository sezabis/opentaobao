package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTemplateReportAPIRequest 服务商存量短信模板上传 API请求
// taobao.jst.sms.template.report
//
// 用于上传目前已经在阿里通信申请到的且正在使用的模板信息，确保模板数据正确，否则会导致短信发送失败！！！
type TaobaoJstSmsTemplateReportAPIRequest struct {
	model.Params
	// 存量短信模板上报入参
	_smsTemplateRequest *SmsTemplateRequest
}

// NewTaobaoJstSmsTemplateReportRequest 初始化TaobaoJstSmsTemplateReportAPIRequest对象
func NewTaobaoJstSmsTemplateReportRequest() *TaobaoJstSmsTemplateReportAPIRequest {
	return &TaobaoJstSmsTemplateReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTemplateReportAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.template.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTemplateReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSmsTemplateRequest is SmsTemplateRequest Setter
// 存量短信模板上报入参
func (r *TaobaoJstSmsTemplateReportAPIRequest) SetSmsTemplateRequest(_smsTemplateRequest *SmsTemplateRequest) error {
	r._smsTemplateRequest = _smsTemplateRequest
	r.Set("sms_template_request", _smsTemplateRequest)
	return nil
}

// GetSmsTemplateRequest SmsTemplateRequest Getter
func (r TaobaoJstSmsTemplateReportAPIRequest) GetSmsTemplateRequest() *SmsTemplateRequest {
	return r._smsTemplateRequest
}
