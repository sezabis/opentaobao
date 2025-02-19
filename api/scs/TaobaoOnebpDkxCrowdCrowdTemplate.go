package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCrowdCrowdTemplate 获取人群模版
// taobao.onebp.dkx.crowd.crowd.template
//
// 获取人群模版，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func TaobaoOnebpDkxCrowdCrowdTemplate(clt *core.SDKClient, req *scs.TaobaoOnebpDkxCrowdCrowdTemplateAPIRequest, session string) (*scs.TaobaoOnebpDkxCrowdCrowdTemplateAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxCrowdCrowdTemplateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
