package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoTopBizSellerSign 淘宝订单履约-商家erp签约
// taobao.top.biz.seller.sign
//
// 淘宝订单履约-商家erp签约，包含各场景的签约
func TaobaoTopBizSellerSign(clt *core.SDKClient, req *jst.TaobaoTopBizSellerSignAPIRequest, session string) (*jst.TaobaoTopBizSellerSignAPIResponse, error) {
	var resp jst.TaobaoTopBizSellerSignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
