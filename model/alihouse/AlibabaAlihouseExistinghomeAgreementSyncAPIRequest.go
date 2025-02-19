package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeAgreementSyncAPIRequest 二手房电子协议数据同步 API请求
// alibaba.alihouse.existinghome.agreement.sync
//
// 二手房电子协议数据同步
type AlibabaAlihouseExistinghomeAgreementSyncAPIRequest struct {
	model.Params
	// 数据结构
	_existingHomeElectricAgreementDto *ExistingHomeElectricAgreementDto
}

// NewAlibabaAlihouseExistinghomeAgreementSyncRequest 初始化AlibabaAlihouseExistinghomeAgreementSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeAgreementSyncRequest() *AlibabaAlihouseExistinghomeAgreementSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeAgreementSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.agreement.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExistingHomeElectricAgreementDto is ExistingHomeElectricAgreementDto Setter
// 数据结构
func (r *AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) SetExistingHomeElectricAgreementDto(_existingHomeElectricAgreementDto *ExistingHomeElectricAgreementDto) error {
	r._existingHomeElectricAgreementDto = _existingHomeElectricAgreementDto
	r.Set("existing_home_electric_agreement_dto", _existingHomeElectricAgreementDto)
	return nil
}

// GetExistingHomeElectricAgreementDto ExistingHomeElectricAgreementDto Getter
func (r AlibabaAlihouseExistinghomeAgreementSyncAPIRequest) GetExistingHomeElectricAgreementDto() *ExistingHomeElectricAgreementDto {
	return r._existingHomeElectricAgreementDto
}
