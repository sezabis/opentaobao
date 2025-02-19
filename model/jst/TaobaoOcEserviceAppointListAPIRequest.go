package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcEserviceAppointListAPIRequest 交互卡片预约信息读取接口 API请求
// taobao.oc.eservice.appoint.list
//
// 允许外部的isv通过该接口读取门店预约信息
type TaobaoOcEserviceAppointListAPIRequest struct {
	model.Params
	// 买家客户nick(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
	_customerNick string
	// 买家客户电话号码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
	_customerPhone string
	// 买家客户装修房屋所在的市(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
	_houseAddressCity string
	// 门店编码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
	_mallCode string
	// 返回结果按预约时间排序，指示升序还是降息，取值asc和desc
	_sortOrder string
	// 查询预约的起始时间，格式yyyyMMddHHmmss，默认为当前时间
	_startAppointTime string
	// 预约信息唯一编码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
	_code int64
	// 卖家主账号id
	_sellerId int64
}

// NewTaobaoOcEserviceAppointListRequest 初始化TaobaoOcEserviceAppointListAPIRequest对象
func NewTaobaoOcEserviceAppointListRequest() *TaobaoOcEserviceAppointListAPIRequest {
	return &TaobaoOcEserviceAppointListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOcEserviceAppointListAPIRequest) GetApiMethodName() string {
	return "taobao.oc.eservice.appoint.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOcEserviceAppointListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCustomerNick is CustomerNick Setter
// 买家客户nick(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
func (r *TaobaoOcEserviceAppointListAPIRequest) SetCustomerNick(_customerNick string) error {
	r._customerNick = _customerNick
	r.Set("customer_nick", _customerNick)
	return nil
}

// GetCustomerNick CustomerNick Getter
func (r TaobaoOcEserviceAppointListAPIRequest) GetCustomerNick() string {
	return r._customerNick
}

// SetCustomerPhone is CustomerPhone Setter
// 买家客户电话号码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
func (r *TaobaoOcEserviceAppointListAPIRequest) SetCustomerPhone(_customerPhone string) error {
	r._customerPhone = _customerPhone
	r.Set("customer_phone", _customerPhone)
	return nil
}

// GetCustomerPhone CustomerPhone Getter
func (r TaobaoOcEserviceAppointListAPIRequest) GetCustomerPhone() string {
	return r._customerPhone
}

// SetHouseAddressCity is HouseAddressCity Setter
// 买家客户装修房屋所在的市(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
func (r *TaobaoOcEserviceAppointListAPIRequest) SetHouseAddressCity(_houseAddressCity string) error {
	r._houseAddressCity = _houseAddressCity
	r.Set("house_address_city", _houseAddressCity)
	return nil
}

// GetHouseAddressCity HouseAddressCity Getter
func (r TaobaoOcEserviceAppointListAPIRequest) GetHouseAddressCity() string {
	return r._houseAddressCity
}

// SetMallCode is MallCode Setter
// 门店编码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
func (r *TaobaoOcEserviceAppointListAPIRequest) SetMallCode(_mallCode string) error {
	r._mallCode = _mallCode
	r.Set("mall_code", _mallCode)
	return nil
}

// GetMallCode MallCode Getter
func (r TaobaoOcEserviceAppointListAPIRequest) GetMallCode() string {
	return r._mallCode
}

// SetSortOrder is SortOrder Setter
// 返回结果按预约时间排序，指示升序还是降息，取值asc和desc
func (r *TaobaoOcEserviceAppointListAPIRequest) SetSortOrder(_sortOrder string) error {
	r._sortOrder = _sortOrder
	r.Set("sort_order", _sortOrder)
	return nil
}

// GetSortOrder SortOrder Getter
func (r TaobaoOcEserviceAppointListAPIRequest) GetSortOrder() string {
	return r._sortOrder
}

// SetStartAppointTime is StartAppointTime Setter
// 查询预约的起始时间，格式yyyyMMddHHmmss，默认为当前时间
func (r *TaobaoOcEserviceAppointListAPIRequest) SetStartAppointTime(_startAppointTime string) error {
	r._startAppointTime = _startAppointTime
	r.Set("start_appoint_time", _startAppointTime)
	return nil
}

// GetStartAppointTime StartAppointTime Getter
func (r TaobaoOcEserviceAppointListAPIRequest) GetStartAppointTime() string {
	return r._startAppointTime
}

// SetCode is Code Setter
// 预约信息唯一编码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
func (r *TaobaoOcEserviceAppointListAPIRequest) SetCode(_code int64) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoOcEserviceAppointListAPIRequest) GetCode() int64 {
	return r._code
}

// SetSellerId is SellerId Setter
// 卖家主账号id
func (r *TaobaoOcEserviceAppointListAPIRequest) SetSellerId(_sellerId int64) error {
	r._sellerId = _sellerId
	r.Set("seller_id", _sellerId)
	return nil
}

// GetSellerId SellerId Getter
func (r TaobaoOcEserviceAppointListAPIRequest) GetSellerId() int64 {
	return r._sellerId
}
