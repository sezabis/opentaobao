package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserOpenuidGetbynickAPIRequest 根据买家nick获取买家openuid API请求
// taobao.user.openuid.getbynick
//
// 根据买家nick获取买家openuid，最大查询30个
type TaobaoUserOpenuidGetbynickAPIRequest struct {
	model.Params
	// 买家nick列表
	_buyerNicks string
}

// NewTaobaoUserOpenuidGetbynickRequest 初始化TaobaoUserOpenuidGetbynickAPIRequest对象
func NewTaobaoUserOpenuidGetbynickRequest() *TaobaoUserOpenuidGetbynickAPIRequest {
	return &TaobaoUserOpenuidGetbynickAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUserOpenuidGetbynickAPIRequest) GetApiMethodName() string {
	return "taobao.user.openuid.getbynick"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUserOpenuidGetbynickAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBuyerNicks is BuyerNicks Setter
// 买家nick列表
func (r *TaobaoUserOpenuidGetbynickAPIRequest) SetBuyerNicks(_buyerNicks string) error {
	r._buyerNicks = _buyerNicks
	r.Set("buyer_nicks", _buyerNicks)
	return nil
}

// GetBuyerNicks BuyerNicks Getter
func (r TaobaoUserOpenuidGetbynickAPIRequest) GetBuyerNicks() string {
	return r._buyerNicks
}
