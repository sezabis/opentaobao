package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyProcessAPIRequest 政策进度查询 API请求
// alitrip.policy.process
//
// 上传特殊类型的单程/往返政策
type AlitripPolicyProcessAPIRequest struct {
	model.Params
	// 进度请求入参
	_policyTaskQueryDTO *PolicyTaskQueryDto
}

// NewAlitripPolicyProcessRequest 初始化AlitripPolicyProcessAPIRequest对象
func NewAlitripPolicyProcessRequest() *AlitripPolicyProcessAPIRequest {
	return &AlitripPolicyProcessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPolicyProcessAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPolicyProcessAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPolicyTaskQueryDTO is PolicyTaskQueryDTO Setter
// 进度请求入参
func (r *AlitripPolicyProcessAPIRequest) SetPolicyTaskQueryDTO(_policyTaskQueryDTO *PolicyTaskQueryDto) error {
	r._policyTaskQueryDTO = _policyTaskQueryDTO
	r.Set("policy_task_query_d_t_o", _policyTaskQueryDTO)
	return nil
}

// GetPolicyTaskQueryDTO PolicyTaskQueryDTO Getter
func (r AlitripPolicyProcessAPIRequest) GetPolicyTaskQueryDTO() *PolicyTaskQueryDto {
	return r._policyTaskQueryDTO
}
