package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbroomtypeDeleteAPIRequest 民宿房源删除接口 API请求
// taobao.xhotel.bnbroomtype.delete
//
// 增加民宿房源删除接口
type TaobaoXhotelBnbroomtypeDeleteAPIRequest struct {
	model.Params
	// 房型Id，系统商outer_id
	_outerId string
	// 对接系统商名称：可为空不要乱填，需要申请后使用
	_vendor string
	// 房型Id，传参方式为rid或outer_id+vendor
	_rid int64
}

// NewTaobaoXhotelBnbroomtypeDeleteRequest 初始化TaobaoXhotelBnbroomtypeDeleteAPIRequest对象
func NewTaobaoXhotelBnbroomtypeDeleteRequest() *TaobaoXhotelBnbroomtypeDeleteAPIRequest {
	return &TaobaoXhotelBnbroomtypeDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbroomtypeDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbroomtype.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbroomtypeDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterId is OuterId Setter
// 房型Id，系统商outer_id
func (r *TaobaoXhotelBnbroomtypeDeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelBnbroomtypeDeleteAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetVendor is Vendor Setter
// 对接系统商名称：可为空不要乱填，需要申请后使用
func (r *TaobaoXhotelBnbroomtypeDeleteAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelBnbroomtypeDeleteAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRid is Rid Setter
// 房型Id，传参方式为rid或outer_id+vendor
func (r *TaobaoXhotelBnbroomtypeDeleteAPIRequest) SetRid(_rid int64) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// GetRid Rid Getter
func (r TaobaoXhotelBnbroomtypeDeleteAPIRequest) GetRid() int64 {
	return r._rid
}
