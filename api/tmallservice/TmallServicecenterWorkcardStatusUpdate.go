package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardStatusUpdate 服务商反馈服务的执行情况
// tmall.servicecenter.workcard.status.update
//
// 1 如果服务商受理了此服务，修改合同状态为：已受理=3&lt;br/&gt;2 如果服务商没有受理此服务，修改合同状态为：已拒绝=10&lt;br/&gt;3 如果服务商执行了此服务，修改合同状态为：已执行=4&lt;br/&gt;4 如果服务商执行服务成功，修改合同状态为：已完成=5&lt;br/&gt;5 如果此工单为合同类型的工单，当服务商受理了此服务后，会进行分账
func TmallServicecenterWorkcardStatusUpdate(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardStatusUpdateAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardStatusUpdateAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardStatusUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
