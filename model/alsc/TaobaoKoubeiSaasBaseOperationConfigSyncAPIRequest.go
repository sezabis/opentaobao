package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest 商家基础经营设置信息同步 API请求
// taobao.koubei.saas.base.operation.config.sync
//
// ISV接入口碑SAAS后, 经营设置数据同步到口碑SAAS
type TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest struct {
	model.Params
	// 商户ID
	_merchantId string
	// 请求ID
	_requestId string
	// 操作员ID
	_outerOperatorId string
	// 业务类型。支付方式：payment_method
	_bizType string
	// 经营设置json串
	_operationConfig string
}

// NewTaobaoKoubeiSaasBaseOperationConfigSyncRequest 初始化TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest对象
func NewTaobaoKoubeiSaasBaseOperationConfigSyncRequest() *TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest {
	return &TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.saas.base.operation.config.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMerchantId is MerchantId Setter
// 商户ID
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) SetMerchantId(_merchantId string) error {
	r._merchantId = _merchantId
	r.Set("merchant_id", _merchantId)
	return nil
}

// GetMerchantId MerchantId Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) GetMerchantId() string {
	return r._merchantId
}

// SetRequestId is RequestId Setter
// 请求ID
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetOuterOperatorId is OuterOperatorId Setter
// 操作员ID
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) SetOuterOperatorId(_outerOperatorId string) error {
	r._outerOperatorId = _outerOperatorId
	r.Set("outer_operator_id", _outerOperatorId)
	return nil
}

// GetOuterOperatorId OuterOperatorId Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) GetOuterOperatorId() string {
	return r._outerOperatorId
}

// SetBizType is BizType Setter
// 业务类型。支付方式：payment_method
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) GetBizType() string {
	return r._bizType
}

// SetOperationConfig is OperationConfig Setter
// 经营设置json串
func (r *TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) SetOperationConfig(_operationConfig string) error {
	r._operationConfig = _operationConfig
	r.Set("operation_config", _operationConfig)
	return nil
}

// GetOperationConfig OperationConfig Getter
func (r TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest) GetOperationConfig() string {
	return r._operationConfig
}
