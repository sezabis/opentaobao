package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCreativeCreativeReportRealtime 获取创意实时报表
// taobao.onebp.dkx.creative.creative.report.realtime
//
// 获取创意实时报表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
func TaobaoOnebpDkxCreativeCreativeReportRealtime(clt *core.SDKClient, req *scs.TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIRequest, session string) (*scs.TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
