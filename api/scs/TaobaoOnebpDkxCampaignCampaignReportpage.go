package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCampaignCampaignReportpage 获取场景计划的报表数据
// taobao.onebp.dkx.campaign.campaign.reportpage
//
// 获取场景计划的报表数据。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func TaobaoOnebpDkxCampaignCampaignReportpage(clt *core.SDKClient, req *scs.TaobaoOnebpDkxCampaignCampaignReportpageAPIRequest, session string) (*scs.TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
