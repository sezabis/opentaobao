package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundNewdetail 商旅机票退票详情接口V2
// alitrip.btrip.flight.distribution.refund.newdetail
//
// 商旅机票退票详情接口V2
func AlitripBtripFlightDistributionRefundNewdetail(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundNewdetailAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionRefundNewdetailAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionRefundNewdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
