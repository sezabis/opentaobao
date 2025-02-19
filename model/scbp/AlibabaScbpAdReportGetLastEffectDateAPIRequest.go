package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportGetLastEffectDateAPIRequest 获取最近报表生成时间 API请求
// alibaba.scbp.ad.report.get.last.effect.date
//
// 获取最近报表生成时间
type AlibabaScbpAdReportGetLastEffectDateAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdReportGetLastEffectDateRequest 初始化AlibabaScbpAdReportGetLastEffectDateAPIRequest对象
func NewAlibabaScbpAdReportGetLastEffectDateRequest() *AlibabaScbpAdReportGetLastEffectDateAPIRequest {
	return &AlibabaScbpAdReportGetLastEffectDateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetLastEffectDateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.get.last.effect.date"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetLastEffectDateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetLastEffectDateAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdReportGetLastEffectDateAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
