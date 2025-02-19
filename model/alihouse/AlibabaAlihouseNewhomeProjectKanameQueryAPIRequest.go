package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest 查询KA楼盘名称 API请求
// alibaba.alihouse.newhome.project.kaname.query
//
// 查询KA楼盘名称
type AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest struct {
	model.Params
	// KA楼盘ID
	_kaProjectMid int64
}

// NewAlibabaAlihouseNewhomeProjectKanameQueryRequest 初始化AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectKanameQueryRequest() *AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest {
	return &AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.kaname.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetKaProjectMid is KaProjectMid Setter
// KA楼盘ID
func (r *AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) SetKaProjectMid(_kaProjectMid int64) error {
	r._kaProjectMid = _kaProjectMid
	r.Set("ka_project_mid", _kaProjectMid)
	return nil
}

// GetKaProjectMid KaProjectMid Getter
func (r AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) GetKaProjectMid() int64 {
	return r._kaProjectMid
}
