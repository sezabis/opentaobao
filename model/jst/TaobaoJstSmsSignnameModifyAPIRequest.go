package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameModifyAPIRequest 淘宝短信签名修改 API请求
// taobao.jst.sms.signname.modify
//
// 淘宝短信签名修改，只能修改还未被审核的签名。
type TaobaoJstSmsSignnameModifyAPIRequest struct {
	model.Params
	// 修改签名入参
	_modifySmsSignRequest *TopModifySmsSignRequest
}

// NewTaobaoJstSmsSignnameModifyRequest 初始化TaobaoJstSmsSignnameModifyAPIRequest对象
func NewTaobaoJstSmsSignnameModifyRequest() *TaobaoJstSmsSignnameModifyAPIRequest {
	return &TaobaoJstSmsSignnameModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsSignnameModifyAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsSignnameModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetModifySmsSignRequest is ModifySmsSignRequest Setter
// 修改签名入参
func (r *TaobaoJstSmsSignnameModifyAPIRequest) SetModifySmsSignRequest(_modifySmsSignRequest *TopModifySmsSignRequest) error {
	r._modifySmsSignRequest = _modifySmsSignRequest
	r.Set("modify_sms_sign_request", _modifySmsSignRequest)
	return nil
}

// GetModifySmsSignRequest ModifySmsSignRequest Getter
func (r TaobaoJstSmsSignnameModifyAPIRequest) GetModifySmsSignRequest() *TopModifySmsSignRequest {
	return r._modifySmsSignRequest
}
