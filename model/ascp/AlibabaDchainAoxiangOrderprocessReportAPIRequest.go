package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangOrderprocessReportAPIRequest 回传仓内作业节点 API请求
// alibaba.dchain.aoxiang.orderprocess.report
//
// 回传仓内作业节点
type AlibabaDchainAoxiangOrderprocessReportAPIRequest struct {
	model.Params
	// 单据作业状态回传参数
	_orderprocessReportRequest *OrderProcessReportRequest
}

// NewAlibabaDchainAoxiangOrderprocessReportRequest 初始化AlibabaDchainAoxiangOrderprocessReportAPIRequest对象
func NewAlibabaDchainAoxiangOrderprocessReportRequest() *AlibabaDchainAoxiangOrderprocessReportAPIRequest {
	return &AlibabaDchainAoxiangOrderprocessReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangOrderprocessReportAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.orderprocess.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangOrderprocessReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderprocessReportRequest is OrderprocessReportRequest Setter
// 单据作业状态回传参数
func (r *AlibabaDchainAoxiangOrderprocessReportAPIRequest) SetOrderprocessReportRequest(_orderprocessReportRequest *OrderProcessReportRequest) error {
	r._orderprocessReportRequest = _orderprocessReportRequest
	r.Set("orderprocess_report_request", _orderprocessReportRequest)
	return nil
}

// GetOrderprocessReportRequest OrderprocessReportRequest Getter
func (r AlibabaDchainAoxiangOrderprocessReportAPIRequest) GetOrderprocessReportRequest() *OrderProcessReportRequest {
	return r._orderprocessReportRequest
}
