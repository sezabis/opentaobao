package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalSolutionServiceResourceQueryAPIRequest 查询解决方案服务资源列表 API请求
// cainiao.global.solution.service.resource.query
//
// 返回直接解决方案的指定物流服务的可用资源列表
type CainiaoGlobalSolutionServiceResourceQueryAPIRequest struct {
	model.Params
	// 多语言信息
	_locale string
	// 查询参数
	_solutionServiceResParam *QuerySolutionServiceResParam
	// 发件信息
	_senderParam *OpenSenderParam
}

// NewCainiaoGlobalSolutionServiceResourceQueryRequest 初始化CainiaoGlobalSolutionServiceResourceQueryAPIRequest对象
func NewCainiaoGlobalSolutionServiceResourceQueryRequest() *CainiaoGlobalSolutionServiceResourceQueryAPIRequest {
	return &CainiaoGlobalSolutionServiceResourceQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalSolutionServiceResourceQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.global.solution.service.resource.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalSolutionServiceResourceQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLocale is Locale Setter
// 多语言信息
func (r *CainiaoGlobalSolutionServiceResourceQueryAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoGlobalSolutionServiceResourceQueryAPIRequest) GetLocale() string {
	return r._locale
}

// SetSolutionServiceResParam is SolutionServiceResParam Setter
// 查询参数
func (r *CainiaoGlobalSolutionServiceResourceQueryAPIRequest) SetSolutionServiceResParam(_solutionServiceResParam *QuerySolutionServiceResParam) error {
	r._solutionServiceResParam = _solutionServiceResParam
	r.Set("solution_service_res_param", _solutionServiceResParam)
	return nil
}

// GetSolutionServiceResParam SolutionServiceResParam Getter
func (r CainiaoGlobalSolutionServiceResourceQueryAPIRequest) GetSolutionServiceResParam() *QuerySolutionServiceResParam {
	return r._solutionServiceResParam
}

// SetSenderParam is SenderParam Setter
// 发件信息
func (r *CainiaoGlobalSolutionServiceResourceQueryAPIRequest) SetSenderParam(_senderParam *OpenSenderParam) error {
	r._senderParam = _senderParam
	r.Set("sender_param", _senderParam)
	return nil
}

// GetSenderParam SenderParam Getter
func (r CainiaoGlobalSolutionServiceResourceQueryAPIRequest) GetSenderParam() *OpenSenderParam {
	return r._senderParam
}
