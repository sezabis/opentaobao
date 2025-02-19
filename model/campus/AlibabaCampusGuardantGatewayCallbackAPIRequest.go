package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusGuardantGatewayCallbackAPIRequest 人卡关系回调 API请求
// alibaba.campus.guardant.gateway.callback
//
// 门禁供应商回调平台通知同步结果
type AlibabaCampusGuardantGatewayCallbackAPIRequest struct {
	model.Params
	// md5
	_token string
	// 请求数据
	_data string
}

// NewAlibabaCampusGuardantGatewayCallbackRequest 初始化AlibabaCampusGuardantGatewayCallbackAPIRequest对象
func NewAlibabaCampusGuardantGatewayCallbackRequest() *AlibabaCampusGuardantGatewayCallbackAPIRequest {
	return &AlibabaCampusGuardantGatewayCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusGuardantGatewayCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guardant.gateway.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusGuardantGatewayCallbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetToken is Token Setter
// md5
func (r *AlibabaCampusGuardantGatewayCallbackAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaCampusGuardantGatewayCallbackAPIRequest) GetToken() string {
	return r._token
}

// SetData is Data Setter
// 请求数据
func (r *AlibabaCampusGuardantGatewayCallbackAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaCampusGuardantGatewayCallbackAPIRequest) GetData() string {
	return r._data
}
