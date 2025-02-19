package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitSealPushAPIRequest 法宝推送用印 API请求
// alibaba.legal.suit.seal.push
//
// 法宝推送用印
type AlibabaLegalSuitSealPushAPIRequest struct {
	model.Params
	// 最外层list
	_sealTaskModel *SealTaskModel
}

// NewAlibabaLegalSuitSealPushRequest 初始化AlibabaLegalSuitSealPushAPIRequest对象
func NewAlibabaLegalSuitSealPushRequest() *AlibabaLegalSuitSealPushAPIRequest {
	return &AlibabaLegalSuitSealPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitSealPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.seal.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitSealPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSealTaskModel is SealTaskModel Setter
// 最外层list
func (r *AlibabaLegalSuitSealPushAPIRequest) SetSealTaskModel(_sealTaskModel *SealTaskModel) error {
	r._sealTaskModel = _sealTaskModel
	r.Set("seal_task_model", _sealTaskModel)
	return nil
}

// GetSealTaskModel SealTaskModel Getter
func (r AlibabaLegalSuitSealPushAPIRequest) GetSealTaskModel() *SealTaskModel {
	return r._sealTaskModel
}
