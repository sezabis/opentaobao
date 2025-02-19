package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameReportAPIRequest 服务商存量短信签名上传 API请求
// taobao.jst.sms.signname.report
//
// 用于上传目前已经在阿里通信申请到的且正在使用的签名数据，确保签名数据正确，否则会导致短信发送失败！！！
type TaobaoJstSmsSignnameReportAPIRequest struct {
	model.Params
	// 上传签名入参
	_smsSignNameRequest *SmsSignNameRequest
}

// NewTaobaoJstSmsSignnameReportRequest 初始化TaobaoJstSmsSignnameReportAPIRequest对象
func NewTaobaoJstSmsSignnameReportRequest() *TaobaoJstSmsSignnameReportAPIRequest {
	return &TaobaoJstSmsSignnameReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsSignnameReportAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsSignnameReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSmsSignNameRequest is SmsSignNameRequest Setter
// 上传签名入参
func (r *TaobaoJstSmsSignnameReportAPIRequest) SetSmsSignNameRequest(_smsSignNameRequest *SmsSignNameRequest) error {
	r._smsSignNameRequest = _smsSignNameRequest
	r.Set("sms_sign_name_request", _smsSignNameRequest)
	return nil
}

// GetSmsSignNameRequest SmsSignNameRequest Getter
func (r TaobaoJstSmsSignnameReportAPIRequest) GetSmsSignNameRequest() *SmsSignNameRequest {
	return r._smsSignNameRequest
}
