package tmallcampus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCampusAuthstatusQueryAPIRequest 学生认证状态查询 API请求
// tmall.campus.authstatus.query
//
// 学生认证状态查询
type TmallCampusAuthstatusQueryAPIRequest struct {
	model.Params
}

// NewTmallCampusAuthstatusQueryRequest 初始化TmallCampusAuthstatusQueryAPIRequest对象
func NewTmallCampusAuthstatusQueryRequest() *TmallCampusAuthstatusQueryAPIRequest {
	return &TmallCampusAuthstatusQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCampusAuthstatusQueryAPIRequest) GetApiMethodName() string {
	return "tmall.campus.authstatus.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCampusAuthstatusQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
