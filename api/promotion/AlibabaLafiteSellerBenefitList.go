package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaLafiteSellerBenefitList 商家自运营权益列表
// alibaba.lafite.seller.benefit.list
//
// 小程序isv可使用该接口获取权益列表
func AlibabaLafiteSellerBenefitList(clt *core.SDKClient, req *promotion.AlibabaLafiteSellerBenefitListAPIRequest, session string) (*promotion.AlibabaLafiteSellerBenefitListAPIResponse, error) {
	var resp promotion.AlibabaLafiteSellerBenefitListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
