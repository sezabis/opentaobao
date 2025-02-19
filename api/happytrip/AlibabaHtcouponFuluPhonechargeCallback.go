package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHtcouponFuluPhonechargeCallback 话费充值回调
// alibaba.htcoupon.fulu.phonecharge.callback
//
// 话费充值为异步操作，此接口用于充值成功后，供应商回调。
func AlibabaHtcouponFuluPhonechargeCallback(clt *core.SDKClient, req *happytrip.AlibabaHtcouponFuluPhonechargeCallbackAPIRequest, session string) (*happytrip.AlibabaHtcouponFuluPhonechargeCallbackAPIResponse, error) {
	var resp happytrip.AlibabaHtcouponFuluPhonechargeCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
