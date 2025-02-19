package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemSkuDelete 删除SKU
// taobao.item.sku.delete
//
// 删除一个sku的数据&lt;br/&gt;需要删除的sku通过属性properties进行匹配查找
func TaobaoItemSkuDelete(clt *core.SDKClient, req *product.TaobaoItemSkuDeleteAPIRequest, session string) (*product.TaobaoItemSkuDeleteAPIResponse, error) {
	var resp product.TaobaoItemSkuDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
