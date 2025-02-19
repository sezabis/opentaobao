package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelPriceGetAPIRequest 酒店报价服务-阿信 API请求
// taobao.alitrip.travel.axin.hotel.price.get
//
// 酒店报价查询服务
type TaobaoAlitripTravelAxinHotelPriceGetAPIRequest struct {
	model.Params
	// 入住日期
	_checkIn string
	// 离店日期
	_checkOut string
	// 标准酒店id
	_shid int64
	// 分销商ID
	_distributorTid int64
	// 城市代码
	_cityCode int64
}

// NewTaobaoAlitripTravelAxinHotelPriceGetRequest 初始化TaobaoAlitripTravelAxinHotelPriceGetAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelPriceGetRequest() *TaobaoAlitripTravelAxinHotelPriceGetAPIRequest {
	return &TaobaoAlitripTravelAxinHotelPriceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCheckIn is CheckIn Setter
// 入住日期
func (r *TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) SetCheckIn(_checkIn string) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) GetCheckIn() string {
	return r._checkIn
}

// SetCheckOut is CheckOut Setter
// 离店日期
func (r *TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetShid is Shid Setter
// 标准酒店id
func (r *TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) GetShid() int64 {
	return r._shid
}

// SetDistributorTid is DistributorTid Setter
// 分销商ID
func (r *TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetCityCode is CityCode Setter
// 城市代码
func (r *TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoAlitripTravelAxinHotelPriceGetAPIRequest) GetCityCode() int64 {
	return r._cityCode
}
