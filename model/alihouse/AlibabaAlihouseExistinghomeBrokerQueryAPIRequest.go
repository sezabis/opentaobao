package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrokerQueryAPIRequest 根据外部经纪人ID查询 API请求
// alibaba.alihouse.existinghome.broker.query
//
// 根据外部经纪人ID查询
type AlibabaAlihouseExistinghomeBrokerQueryAPIRequest struct {
	model.Params
	// 外部经纪人ID
	_outerBrokerId string
}

// NewAlibabaAlihouseExistinghomeBrokerQueryRequest 初始化AlibabaAlihouseExistinghomeBrokerQueryAPIRequest对象
func NewAlibabaAlihouseExistinghomeBrokerQueryRequest() *AlibabaAlihouseExistinghomeBrokerQueryAPIRequest {
	return &AlibabaAlihouseExistinghomeBrokerQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.broker.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterBrokerId is OuterBrokerId Setter
// 外部经纪人ID
func (r *AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) SetOuterBrokerId(_outerBrokerId string) error {
	r._outerBrokerId = _outerBrokerId
	r.Set("outer_broker_id", _outerBrokerId)
	return nil
}

// GetOuterBrokerId OuterBrokerId Getter
func (r AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) GetOuterBrokerId() string {
	return r._outerBrokerId
}
