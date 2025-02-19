package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest 代理商获取退票详情回调 API请求
// taobao.train.agent.returnticketinfo.get.vtwo
//
// 代理商获取退票详情回调
type TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
	// 淘宝子订单号
	_subOrderId int64
	// 淘宝主订单号
	_mainOrderId int64
}

// NewTaobaoTrainAgentReturnticketinfoGetVtwoRequest 初始化TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest对象
func NewTaobaoTrainAgentReturnticketinfoGetVtwoRequest() *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest {
	return &TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.returnticketinfo.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetSubOrderId is SubOrderId Setter
// 淘宝子订单号
func (r *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}

// SetMainOrderId is MainOrderId Setter
// 淘宝主订单号
func (r *TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoTrainAgentReturnticketinfoGetVtwoAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}
