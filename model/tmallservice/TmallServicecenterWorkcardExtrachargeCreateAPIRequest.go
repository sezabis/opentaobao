package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardExtrachargeCreateAPIRequest 创建工单额外收费项 API请求
// tmall.servicecenter.workcard.extracharge.create
//
// 创建额外收费项
type TmallServicecenterWorkcardExtrachargeCreateAPIRequest struct {
	model.Params
	// 工单id
	_workcardId int64
	// 额外收费项列表
	_extraChargeItemList *WorkcardExtraChargeCreateTuple
}

// NewTmallServicecenterWorkcardExtrachargeCreateRequest 初始化TmallServicecenterWorkcardExtrachargeCreateAPIRequest对象
func NewTmallServicecenterWorkcardExtrachargeCreateRequest() *TmallServicecenterWorkcardExtrachargeCreateAPIRequest {
	return &TmallServicecenterWorkcardExtrachargeCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardExtrachargeCreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.extracharge.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardExtrachargeCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardExtrachargeCreateAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterWorkcardExtrachargeCreateAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetExtraChargeItemList is ExtraChargeItemList Setter
// 额外收费项列表
func (r *TmallServicecenterWorkcardExtrachargeCreateAPIRequest) SetExtraChargeItemList(_extraChargeItemList *WorkcardExtraChargeCreateTuple) error {
	r._extraChargeItemList = _extraChargeItemList
	r.Set("extra_charge_item_list", _extraChargeItemList)
	return nil
}

// GetExtraChargeItemList ExtraChargeItemList Getter
func (r TmallServicecenterWorkcardExtrachargeCreateAPIRequest) GetExtraChargeItemList() *WorkcardExtraChargeCreateTuple {
	return r._extraChargeItemList
}
