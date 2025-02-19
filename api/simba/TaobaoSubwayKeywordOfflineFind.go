package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayKeywordOfflineFind 查询关键词离线多日汇总报表
// taobao.subway.keyword.offline.find
//
// 获取关键词离线报表
func TaobaoSubwayKeywordOfflineFind(clt *core.SDKClient, req *simba.TaobaoSubwayKeywordOfflineFindAPIRequest, session string) (*simba.TaobaoSubwayKeywordOfflineFindAPIResponse, error) {
	var resp simba.TaobaoSubwayKeywordOfflineFindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
