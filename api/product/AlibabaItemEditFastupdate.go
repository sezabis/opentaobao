package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaItemEditFastupdate 商品编辑增量更新
// alibaba.item.edit.fastupdate
//
// 商品编辑增量更新;
// &lt;br/&gt;该接口编辑sku，只能更新价格、库存等信息，不能新增sku;
// &lt;br/&gt;新增sku用全量接口alibaba.item.edit.submit，先设置销售属性;
func AlibabaItemEditFastupdate(clt *core.SDKClient, req *product.AlibabaItemEditFastupdateAPIRequest, session string) (*product.AlibabaItemEditFastupdateAPIResponse, error) {
	var resp product.AlibabaItemEditFastupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
