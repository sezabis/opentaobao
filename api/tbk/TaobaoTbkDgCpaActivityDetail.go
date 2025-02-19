package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgCpaActivityDetail 淘宝客-推广者-CPA活动执行明细
// taobao.tbk.dg.cpa.activity.detail
//
// 淘宝客获取CPA活动具体执行效果的明细数据（含预估和结算维度）
func TaobaoTbkDgCpaActivityDetail(clt *core.SDKClient, req *tbk.TaobaoTbkDgCpaActivityDetailAPIRequest, session string) (*tbk.TaobaoTbkDgCpaActivityDetailAPIResponse, error) {
	var resp tbk.TaobaoTbkDgCpaActivityDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
