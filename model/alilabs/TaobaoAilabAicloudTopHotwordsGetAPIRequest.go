package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopHotwordsGetAPIRequest 获取热词 API请求
// taobao.ailab.aicloud.top.hotwords.get
//
// 获取ASR热词
type TaobaoAilabAicloudTopHotwordsGetAPIRequest struct {
	model.Params
	// schemeKey
	_schema string
	// 三方用户id
	_userId string
	// 业务类型
	_bizClass string
}

// NewTaobaoAilabAicloudTopHotwordsGetRequest 初始化TaobaoAilabAicloudTopHotwordsGetAPIRequest对象
func NewTaobaoAilabAicloudTopHotwordsGetRequest() *TaobaoAilabAicloudTopHotwordsGetAPIRequest {
	return &TaobaoAilabAicloudTopHotwordsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopHotwordsGetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.hotwords.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopHotwordsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSchema is Schema Setter
// schemeKey
func (r *TaobaoAilabAicloudTopHotwordsGetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopHotwordsGetAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 三方用户id
func (r *TaobaoAilabAicloudTopHotwordsGetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopHotwordsGetAPIRequest) GetUserId() string {
	return r._userId
}

// SetBizClass is BizClass Setter
// 业务类型
func (r *TaobaoAilabAicloudTopHotwordsGetAPIRequest) SetBizClass(_bizClass string) error {
	r._bizClass = _bizClass
	r.Set("biz_class", _bizClass)
	return nil
}

// GetBizClass BizClass Getter
func (r TaobaoAilabAicloudTopHotwordsGetAPIRequest) GetBizClass() string {
	return r._bizClass
}
