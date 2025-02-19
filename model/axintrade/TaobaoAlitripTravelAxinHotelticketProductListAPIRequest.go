package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketProductListAPIRequest 阿信酒景套餐产品列表查询 API请求
// taobao.alitrip.travel.axin.hotelticket.product.list
//
// 阿信酒景套餐产品列表查询
type TaobaoAlitripTravelAxinHotelticketProductListAPIRequest struct {
	model.Params
	// 分销商id
	_distributorTid int64
	// 当前页数
	_pageNo int64
	// 分页大小
	_pageSize int64
}

// NewTaobaoAlitripTravelAxinHotelticketProductListRequest 初始化TaobaoAlitripTravelAxinHotelticketProductListAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelticketProductListRequest() *TaobaoAlitripTravelAxinHotelticketProductListAPIRequest {
	return &TaobaoAlitripTravelAxinHotelticketProductListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelticketProductListAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotelticket.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelticketProductListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelticketProductListAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelticketProductListAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPageNo is PageNo Setter
// 当前页数
func (r *TaobaoAlitripTravelAxinHotelticketProductListAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoAlitripTravelAxinHotelticketProductListAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaoAlitripTravelAxinHotelticketProductListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoAlitripTravelAxinHotelticketProductListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
