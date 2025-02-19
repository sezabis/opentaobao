package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardDesensitizationQueryAPIRequest 工单查询接口 API请求
// tmall.servicecenter.workcard.desensitization.query
//
// 工单查询接口
type TmallServicecenterWorkcardDesensitizationQueryAPIRequest struct {
	model.Params
	// 核销码
	_identifyCode string
	// 工单创建开始时间
	_gmtCreateStart string
	// 工单创建结束时间，必须与工单创建开始时间一起传入，且间隔不超过15分钟
	_gmtCreateEnd string
	// 门店/网点id
	_serviceStoreId int64
	// 工单id
	_id int64
	// 淘宝交易订单号。主订单或子订单均可
	_bizOrderId int64
	// 当前页数
	_currentPage int64
	// 每页大小
	_pageSize int64
	// 请求来源类型
	_requestSourceType int64
}

// NewTmallServicecenterWorkcardDesensitizationQueryRequest 初始化TmallServicecenterWorkcardDesensitizationQueryAPIRequest对象
func NewTmallServicecenterWorkcardDesensitizationQueryRequest() *TmallServicecenterWorkcardDesensitizationQueryAPIRequest {
	return &TmallServicecenterWorkcardDesensitizationQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.desensitization.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIdentifyCode is IdentifyCode Setter
// 核销码
func (r *TmallServicecenterWorkcardDesensitizationQueryAPIRequest) SetIdentifyCode(_identifyCode string) error {
	r._identifyCode = _identifyCode
	r.Set("identify_code", _identifyCode)
	return nil
}

// GetIdentifyCode IdentifyCode Getter
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetIdentifyCode() string {
	return r._identifyCode
}

// SetGmtCreateStart is GmtCreateStart Setter
// 工单创建开始时间
func (r *TmallServicecenterWorkcardDesensitizationQueryAPIRequest) SetGmtCreateStart(_gmtCreateStart string) error {
	r._gmtCreateStart = _gmtCreateStart
	r.Set("gmt_create_start", _gmtCreateStart)
	return nil
}

// GetGmtCreateStart GmtCreateStart Getter
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetGmtCreateStart() string {
	return r._gmtCreateStart
}

// SetGmtCreateEnd is GmtCreateEnd Setter
// 工单创建结束时间，必须与工单创建开始时间一起传入，且间隔不超过15分钟
func (r *TmallServicecenterWorkcardDesensitizationQueryAPIRequest) SetGmtCreateEnd(_gmtCreateEnd string) error {
	r._gmtCreateEnd = _gmtCreateEnd
	r.Set("gmt_create_end", _gmtCreateEnd)
	return nil
}

// GetGmtCreateEnd GmtCreateEnd Getter
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetGmtCreateEnd() string {
	return r._gmtCreateEnd
}

// SetServiceStoreId is ServiceStoreId Setter
// 门店/网点id
func (r *TmallServicecenterWorkcardDesensitizationQueryAPIRequest) SetServiceStoreId(_serviceStoreId int64) error {
	r._serviceStoreId = _serviceStoreId
	r.Set("service_store_id", _serviceStoreId)
	return nil
}

// GetServiceStoreId ServiceStoreId Getter
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetServiceStoreId() int64 {
	return r._serviceStoreId
}

// SetId is Id Setter
// 工单id
func (r *TmallServicecenterWorkcardDesensitizationQueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetId() int64 {
	return r._id
}

// SetBizOrderId is BizOrderId Setter
// 淘宝交易订单号。主订单或子订单均可
func (r *TmallServicecenterWorkcardDesensitizationQueryAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetCurrentPage is CurrentPage Setter
// 当前页数
func (r *TmallServicecenterWorkcardDesensitizationQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TmallServicecenterWorkcardDesensitizationQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetRequestSourceType is RequestSourceType Setter
// 请求来源类型
func (r *TmallServicecenterWorkcardDesensitizationQueryAPIRequest) SetRequestSourceType(_requestSourceType int64) error {
	r._requestSourceType = _requestSourceType
	r.Set("request_source_type", _requestSourceType)
	return nil
}

// GetRequestSourceType RequestSourceType Getter
func (r TmallServicecenterWorkcardDesensitizationQueryAPIRequest) GetRequestSourceType() int64 {
	return r._requestSourceType
}
