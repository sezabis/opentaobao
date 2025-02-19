package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardRepairprogressUpdateAPIRequest 更新维修进度 API请求
// tmall.servicecenter.workcard.repairprogress.update
//
// 提供给外部合作服务商的维修进度更改接口
type TmallServicecenterWorkcardRepairprogressUpdateAPIRequest struct {
	model.Params
	// 图片列表
	_picUrlList []string
	// 真实接单服务商账号Nick
	_realTpNick string
	// 请求节点的动作描述，唯一标识一个节点
	_action string
	// 衣服，鞋子
	_receivedGoods string
	// 服务目标物瑕疵信息
	_targetGoodsDefects string
	// 工单id
	_workcardId int64
}

// NewTmallServicecenterWorkcardRepairprogressUpdateRequest 初始化TmallServicecenterWorkcardRepairprogressUpdateAPIRequest对象
func NewTmallServicecenterWorkcardRepairprogressUpdateRequest() *TmallServicecenterWorkcardRepairprogressUpdateAPIRequest {
	return &TmallServicecenterWorkcardRepairprogressUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.repairprogress.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPicUrlList is PicUrlList Setter
// 图片列表
func (r *TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) SetPicUrlList(_picUrlList []string) error {
	r._picUrlList = _picUrlList
	r.Set("pic_url_list", _picUrlList)
	return nil
}

// GetPicUrlList PicUrlList Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) GetPicUrlList() []string {
	return r._picUrlList
}

// SetRealTpNick is RealTpNick Setter
// 真实接单服务商账号Nick
func (r *TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetAction is Action Setter
// 请求节点的动作描述，唯一标识一个节点
func (r *TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) GetAction() string {
	return r._action
}

// SetReceivedGoods is ReceivedGoods Setter
// 衣服，鞋子
func (r *TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) SetReceivedGoods(_receivedGoods string) error {
	r._receivedGoods = _receivedGoods
	r.Set("received_goods", _receivedGoods)
	return nil
}

// GetReceivedGoods ReceivedGoods Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) GetReceivedGoods() string {
	return r._receivedGoods
}

// SetTargetGoodsDefects is TargetGoodsDefects Setter
// 服务目标物瑕疵信息
func (r *TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) SetTargetGoodsDefects(_targetGoodsDefects string) error {
	r._targetGoodsDefects = _targetGoodsDefects
	r.Set("target_goods_defects", _targetGoodsDefects)
	return nil
}

// GetTargetGoodsDefects TargetGoodsDefects Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) GetTargetGoodsDefects() string {
	return r._targetGoodsDefects
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterWorkcardRepairprogressUpdateAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}
