package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportCrowdList 获取人群离线报表
// taobao.onebp.dkx.report.report.crowd.list
//
// 获取人群离线报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-08&#34;,&#34;campaign_id_list&#34;:[2821811613],&#34;effect&#34;:15,&#34;end_time&#34;:&#34;2021-09-10&#34;,&#34;crowd_id&#34;:12297883}
func TaobaoOnebpDkxReportReportCrowdList(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportCrowdListAPIRequest, session string) (*scs.TaobaoOnebpDkxReportReportCrowdListAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxReportReportCrowdListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
