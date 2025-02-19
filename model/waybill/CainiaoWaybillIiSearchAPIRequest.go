package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiSearchAPIRequest 查询面单服务订购及面单使用情况 API请求
// cainiao.waybill.ii.search
//
// 获取发货地&amp;CP开通状态&amp;账户的使用情况
type CainiaoWaybillIiSearchAPIRequest struct {
	model.Params
	// 物流公司code
	_cpCode string
}

// NewCainiaoWaybillIiSearchRequest 初始化CainiaoWaybillIiSearchAPIRequest对象
func NewCainiaoWaybillIiSearchRequest() *CainiaoWaybillIiSearchAPIRequest {
	return &CainiaoWaybillIiSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiSearchAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCpCode is CpCode Setter
// 物流公司code
func (r *CainiaoWaybillIiSearchAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaoWaybillIiSearchAPIRequest) GetCpCode() string {
	return r._cpCode
}
