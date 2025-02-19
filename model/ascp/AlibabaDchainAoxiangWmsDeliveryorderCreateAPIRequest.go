package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest 回传仓库接单通知 API请求
// alibaba.dchain.aoxiang.wms.deliveryorder.create
//
// WMS上报仓库接单节点状态信息，代表接单环节。
type AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest struct {
	model.Params
	// 接单回传上报请求
	_deliveryOrderReportRequest *DeliveryOrderReportRequest
}

// NewAlibabaDchainAoxiangWmsDeliveryorderCreateRequest 初始化AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest对象
func NewAlibabaDchainAoxiangWmsDeliveryorderCreateRequest() *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest {
	return &AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.wms.deliveryorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeliveryOrderReportRequest is DeliveryOrderReportRequest Setter
// 接单回传上报请求
func (r *AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) SetDeliveryOrderReportRequest(_deliveryOrderReportRequest *DeliveryOrderReportRequest) error {
	r._deliveryOrderReportRequest = _deliveryOrderReportRequest
	r.Set("delivery_order_report_request", _deliveryOrderReportRequest)
	return nil
}

// GetDeliveryOrderReportRequest DeliveryOrderReportRequest Getter
func (r AlibabaDchainAoxiangWmsDeliveryorderCreateAPIRequest) GetDeliveryOrderReportRequest() *DeliveryOrderReportRequest {
	return r._deliveryOrderReportRequest
}
