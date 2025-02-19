package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformAccountCharge 资金平台余额账户充值
// alibaba.fundplatform.account.charge
//
// 资金平台余额账户充值【创建账户&amp;返回付款URL】
func AlibabaFundplatformAccountCharge(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformAccountChargeAPIRequest, session string) (*fundplatform.AlibabaFundplatformAccountChargeAPIResponse, error) {
	var resp fundplatform.AlibabaFundplatformAccountChargeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
