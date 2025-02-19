package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSpecGetAPIRequest 根据产品规格的Id号获取当个的规格信息 API请求
// tmall.product.spec.get
//
// 通过当个的spec_id获取到这个产品规格的信息，主要是因为产品规格是要经过审核的，所以通过这个接口可以获取到是否通过审核&lt;br/&gt;通过参看这个ProductSpec的status判断：&lt;br/&gt;1:表示审核通过&lt;br/&gt;3:表示等待审核。&lt;br/&gt;如果你的id找不到数据，那么就是审核被拒绝。
type TmallProductSpecGetAPIRequest struct {
	model.Params
	// 要获取信息的产品规格信息。
	_specId int64
}

// NewTmallProductSpecGetRequest 初始化TmallProductSpecGetAPIRequest对象
func NewTmallProductSpecGetRequest() *TmallProductSpecGetAPIRequest {
	return &TmallProductSpecGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductSpecGetAPIRequest) GetApiMethodName() string {
	return "tmall.product.spec.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductSpecGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSpecId is SpecId Setter
// 要获取信息的产品规格信息。
func (r *TmallProductSpecGetAPIRequest) SetSpecId(_specId int64) error {
	r._specId = _specId
	r.Set("spec_id", _specId)
	return nil
}

// GetSpecId SpecId Getter
func (r TmallProductSpecGetAPIRequest) GetSpecId() int64 {
	return r._specId
}
