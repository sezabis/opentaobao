package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIRequest 查询达摩盘精选人群模板 API请求
// taobao.onebp.dkx.crowd.crowd.finddmpcrowd
//
// 查询达摩盘精选人群模板；使用方法为先查询出topic和对应的templateId（需要一一对应），然后将想使用的模板topic&amp;templateId组填入Add接口中的new_dmp_template_crowd结构中提交即可。
type TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
}

// NewTaobaoOnebpDkxCrowdCrowdFinddmpcrowdRequest 初始化TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIRequest对象
func NewTaobaoOnebpDkxCrowdCrowdFinddmpcrowdRequest() *TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIRequest {
	return &TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.crowd.crowd.finddmpcrowd"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxCrowdCrowdFinddmpcrowdAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}
