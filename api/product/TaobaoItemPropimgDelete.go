package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemPropimgDelete 删除属性图片
// taobao.item.propimg.delete
//
// 删除propimg_id 所指定的商品属性图片 &lt;br/&gt;传入的num_iid所对应的商品必须属于当前会话的用户 &lt;br/&gt;propimg_id对应的属性图片需要属于num_iid对应的商品
func TaobaoItemPropimgDelete(clt *core.SDKClient, req *product.TaobaoItemPropimgDeleteAPIRequest, session string) (*product.TaobaoItemPropimgDeleteAPIResponse, error) {
	var resp product.TaobaoItemPropimgDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
