package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportGetTargetReportAPIRequest 定向报告 API请求
// alibaba.scbp.ad.report.get.target.report
//
// 定向报告
type AlibabaScbpAdReportGetTargetReportAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_targetReportOperation *TargetReportOperationDto
}

// NewAlibabaScbpAdReportGetTargetReportRequest 初始化AlibabaScbpAdReportGetTargetReportAPIRequest对象
func NewAlibabaScbpAdReportGetTargetReportRequest() *AlibabaScbpAdReportGetTargetReportAPIRequest {
	return &AlibabaScbpAdReportGetTargetReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetTargetReportAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.get.target.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetTargetReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetTargetReportAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdReportGetTargetReportAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetTargetReportOperation is TargetReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportGetTargetReportAPIRequest) SetTargetReportOperation(_targetReportOperation *TargetReportOperationDto) error {
	r._targetReportOperation = _targetReportOperation
	r.Set("target_report_operation", _targetReportOperation)
	return nil
}

// GetTargetReportOperation TargetReportOperation Getter
func (r AlibabaScbpAdReportGetTargetReportAPIRequest) GetTargetReportOperation() *TargetReportOperationDto {
	return r._targetReportOperation
}
