package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemPropertyDefQuery 交易猫商品属性定义查询
// alibaba.jym.item.property.def.query
//
// 查询商品发布属性定义详情
func AlibabaJymItemPropertyDefQuery(clt *core.SDKClient, req *product.AlibabaJymItemPropertyDefQueryAPIRequest, session string) (*product.AlibabaJymItemPropertyDefQueryAPIResponse, error) {
	var resp product.AlibabaJymItemPropertyDefQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
