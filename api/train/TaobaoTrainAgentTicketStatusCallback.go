package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainAgentTicketStatusCallback 代理商票状态查询回调
// taobao.train.agent.ticket.status.callback
//
// 代理商票状态查询结果回调
func TaobaoTrainAgentTicketStatusCallback(clt *core.SDKClient, req *train.TaobaoTrainAgentTicketStatusCallbackAPIRequest, session string) (*train.TaobaoTrainAgentTicketStatusCallbackAPIResponse, error) {
	var resp train.TaobaoTrainAgentTicketStatusCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
