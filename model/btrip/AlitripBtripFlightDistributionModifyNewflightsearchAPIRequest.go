package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest 改签航班列表V2 API请求
// alitrip.btrip.flight.distribution.modify.newflightsearch
//
// 改签航班列表V2
type AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest struct {
	model.Params
	// 改签航班列表入参
	_paramBtripFlightModifySearchPriceRq *BtripFlightModifySearchPriceRq
}

// NewAlitripBtripFlightDistributionModifyNewflightsearchRequest 初始化AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest对象
func NewAlitripBtripFlightDistributionModifyNewflightsearchRequest() *AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest {
	return &AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.distribution.modify.newflightsearch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBtripFlightModifySearchPriceRq is ParamBtripFlightModifySearchPriceRq Setter
// 改签航班列表入参
func (r *AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest) SetParamBtripFlightModifySearchPriceRq(_paramBtripFlightModifySearchPriceRq *BtripFlightModifySearchPriceRq) error {
	r._paramBtripFlightModifySearchPriceRq = _paramBtripFlightModifySearchPriceRq
	r.Set("param_btrip_flight_modify_search_price_rq", _paramBtripFlightModifySearchPriceRq)
	return nil
}

// GetParamBtripFlightModifySearchPriceRq ParamBtripFlightModifySearchPriceRq Getter
func (r AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest) GetParamBtripFlightModifySearchPriceRq() *BtripFlightModifySearchPriceRq {
	return r._paramBtripFlightModifySearchPriceRq
}
