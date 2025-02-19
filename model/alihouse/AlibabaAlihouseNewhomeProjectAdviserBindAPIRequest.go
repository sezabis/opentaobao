package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest 置业顾问批量绑定楼盘 API请求
// alibaba.alihouse.newhome.project.adviser.bind
//
// 置业顾问批量绑定楼盘
type AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest struct {
	model.Params
	// 对象
	_projectAdviserBindDto *ProjectAdviserBindDto
	// 状态：0无效，1有效
	_status int64
}

// NewAlibabaAlihouseNewhomeProjectAdviserBindRequest 初始化AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectAdviserBindRequest() *AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest {
	return &AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.adviser.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProjectAdviserBindDto is ProjectAdviserBindDto Setter
// 对象
func (r *AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) SetProjectAdviserBindDto(_projectAdviserBindDto *ProjectAdviserBindDto) error {
	r._projectAdviserBindDto = _projectAdviserBindDto
	r.Set("project_adviser_bind_dto", _projectAdviserBindDto)
	return nil
}

// GetProjectAdviserBindDto ProjectAdviserBindDto Getter
func (r AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) GetProjectAdviserBindDto() *ProjectAdviserBindDto {
	return r._projectAdviserBindDto
}

// SetStatus is Status Setter
// 状态：0无效，1有效
func (r *AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAlihouseNewhomeProjectAdviserBindAPIRequest) GetStatus() int64 {
	return r._status
}
