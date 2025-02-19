package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest 取消商品分销 API请求
// alibaba.dchain.aoxiang.item.distribution.batch.cancel
//
// 取消商品分销
type AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest struct {
	model.Params
	// 取消商品分销入参
	_cancelDistributeRequest *CancelDistributeRequest
}

// NewAlibabaDchainAoxiangItemDistributionBatchCancelRequest 初始化AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest对象
func NewAlibabaDchainAoxiangItemDistributionBatchCancelRequest() *AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest {
	return &AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.item.distribution.batch.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCancelDistributeRequest is CancelDistributeRequest Setter
// 取消商品分销入参
func (r *AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) SetCancelDistributeRequest(_cancelDistributeRequest *CancelDistributeRequest) error {
	r._cancelDistributeRequest = _cancelDistributeRequest
	r.Set("cancel_distribute_request", _cancelDistributeRequest)
	return nil
}

// GetCancelDistributeRequest CancelDistributeRequest Getter
func (r AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest) GetCancelDistributeRequest() *CancelDistributeRequest {
	return r._cancelDistributeRequest
}
