package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbownerDeleteAPIRequest 民宿房东删除接口 API请求
// taobao.xhotel.bnbowner.delete
//
// 民宿房东删除接口，删除房东后，对应的门店及房源会同步删除，请谨慎使用
type TaobaoXhotelBnbownerDeleteAPIRequest struct {
	model.Params
	// 对接系统商名称：可为空不要乱填，需要申请后使用，默认taobao
	_vendor string
	// 房东Id，系统商outer_id
	_outerId string
}

// NewTaobaoXhotelBnbownerDeleteRequest 初始化TaobaoXhotelBnbownerDeleteAPIRequest对象
func NewTaobaoXhotelBnbownerDeleteRequest() *TaobaoXhotelBnbownerDeleteAPIRequest {
	return &TaobaoXhotelBnbownerDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbownerDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbowner.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbownerDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetVendor is Vendor Setter
// 对接系统商名称：可为空不要乱填，需要申请后使用，默认taobao
func (r *TaobaoXhotelBnbownerDeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelBnbownerDeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOuterId is OuterId Setter
// 房东Id，系统商outer_id
func (r *TaobaoXhotelBnbownerDeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelBnbownerDeleteAPIRequest) GetOuterId() string {
	return r._outerId
}
