package brandhub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandStarshopRptTargetGetAPIRequest 明星店铺定向维度报表 API请求
// taobao.brand.starshop.rpt.target.get
//
// 获取明星店铺定向维度分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandStarshopRptTargetGetAPIRequest struct {
	model.Params
	// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
	_trafficType string
	// 开始日期(最多查询1个月的数据)
	_startDate string
	// 截至日期(最晚到昨天)
	_endDate string
	// 当前页数
	_pageIndex string
	// 每页条数
	_pageSize string
	// 转化周期,默认15, 3,7,15
	_effect string
}

// NewTaobaoBrandStarshopRptTargetGetRequest 初始化TaobaoBrandStarshopRptTargetGetAPIRequest对象
func NewTaobaoBrandStarshopRptTargetGetRequest() *TaobaoBrandStarshopRptTargetGetAPIRequest {
	return &TaobaoBrandStarshopRptTargetGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBrandStarshopRptTargetGetAPIRequest) GetApiMethodName() string {
	return "taobao.brand.starshop.rpt.target.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBrandStarshopRptTargetGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTrafficType is TrafficType Setter
// 流量类型 1: PC站内, 2: PC站外 , 4: 无线站内, 5: 无线站外,支持多种一起查询,如1,2,4,5
func (r *TaobaoBrandStarshopRptTargetGetAPIRequest) SetTrafficType(_trafficType string) error {
	r._trafficType = _trafficType
	r.Set("traffic_type", _trafficType)
	return nil
}

// GetTrafficType TrafficType Getter
func (r TaobaoBrandStarshopRptTargetGetAPIRequest) GetTrafficType() string {
	return r._trafficType
}

// SetStartDate is StartDate Setter
// 开始日期(最多查询1个月的数据)
func (r *TaobaoBrandStarshopRptTargetGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoBrandStarshopRptTargetGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 截至日期(最晚到昨天)
func (r *TaobaoBrandStarshopRptTargetGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoBrandStarshopRptTargetGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageIndex is PageIndex Setter
// 当前页数
func (r *TaobaoBrandStarshopRptTargetGetAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoBrandStarshopRptTargetGetAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *TaobaoBrandStarshopRptTargetGetAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoBrandStarshopRptTargetGetAPIRequest) GetPageSize() string {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期,默认15, 3,7,15
func (r *TaobaoBrandStarshopRptTargetGetAPIRequest) SetEffect(_effect string) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoBrandStarshopRptTargetGetAPIRequest) GetEffect() string {
	return r._effect
}
