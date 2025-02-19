package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsBatchtaskQuery 交易猫外部商家查询商品批量任务接口
// alibaba.jym.item.external.goods.batchtask.query
//
// 供外部B端商家接入，请求查询商品批量任务，返回商品批量任务查询结果
func AlibabaJymItemExternalGoodsBatchtaskQuery(clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest, session string) (*product.AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse, error) {
	var resp product.AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
