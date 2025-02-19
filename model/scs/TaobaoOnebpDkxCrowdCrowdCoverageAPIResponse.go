package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCrowdCrowdCoverageAPIResponse 获取人数预估 API返回值
// taobao.onebp.dkx.crowd.crowd.coverage
//
// 获取人数预估，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCrowdCrowdCoverageAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCrowdCrowdCoverageAPIResponseModel
}

// TaobaoOnebpDkxCrowdCrowdCoverageAPIResponseModel is 获取人数预估 成功返回结果
type TaobaoOnebpDkxCrowdCrowdCoverageAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_crowd_crowd_coverage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCrowdCrowdCoverageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
