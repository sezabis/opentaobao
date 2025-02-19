package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSkuSortUpdate 商品销售属性排序更新
// tmall.item.sku.sort.update
//
// 商品销售属性排序更新
func TmallItemSkuSortUpdate(clt *core.SDKClient, req *product.TmallItemSkuSortUpdateAPIRequest, session string) (*product.TmallItemSkuSortUpdateAPIResponse, error) {
	var resp product.TmallItemSkuSortUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
