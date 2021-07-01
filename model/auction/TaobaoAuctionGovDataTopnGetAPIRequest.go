package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAuctionGovDataTopnGetAPIRequest
根据不同维度，获取排行榜列表 API请求
taobao.auction.gov.data.topn.get

根据不同时间维度(周,月,年)，获取(成交额或发拍件数)排行榜列表 */
type TaobaoAuctionGovDataTopnGetAPIRequest struct {
	model.Params
	// 周期类型  （2：周，3：月，4：年）
	_circleType int64
	// 周期区间 周期（周填0、月份 yyyyMM、年份 yyyy)
	_circle int64
	// 业务类型 （1：成交额，2：发拍件数）
	_busiType int64
	// 区域类型（1：全国，2：全省）
	_zoneType int64
	// 法院名称
	_courtName string
}

// NewTaobaoAuctionGovDataTopnGetRequest 初始化TaobaoAuctionGovDataTopnGetAPIRequest对象
func NewTaobaoAuctionGovDataTopnGetRequest() *TaobaoAuctionGovDataTopnGetAPIRequest {
	return &TaobaoAuctionGovDataTopnGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetApiMethodName() string {
	return "taobao.auction.gov.data.topn.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CircleType Setter
// 周期类型  （2：周，3：月，4：年）
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) SetCircleType(_circleType int64) error {
	r._circleType = _circleType
	r.Set("circle_type", _circleType)
	return nil
}

// Get CircleType Getter
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetCircleType() int64 {
	return r._circleType
}

// Set is Circle Setter
// 周期区间 周期（周填0、月份 yyyyMM、年份 yyyy)
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) SetCircle(_circle int64) error {
	r._circle = _circle
	r.Set("circle", _circle)
	return nil
}

// Get Circle Getter
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetCircle() int64 {
	return r._circle
}

// Set is BusiType Setter
// 业务类型 （1：成交额，2：发拍件数）
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) SetBusiType(_busiType int64) error {
	r._busiType = _busiType
	r.Set("busi_type", _busiType)
	return nil
}

// Get BusiType Getter
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetBusiType() int64 {
	return r._busiType
}

// Set is ZoneType Setter
// 区域类型（1：全国，2：全省）
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) SetZoneType(_zoneType int64) error {
	r._zoneType = _zoneType
	r.Set("zone_type", _zoneType)
	return nil
}

// Get ZoneType Getter
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetZoneType() int64 {
	return r._zoneType
}

// Set is CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovDataTopnGetAPIRequest) SetCourtName(_courtName string) error {
	r._courtName = _courtName
	r.Set("court_name", _courtName)
	return nil
}

// Get CourtName Getter
func (r TaobaoAuctionGovDataTopnGetAPIRequest) GetCourtName() string {
	return r._courtName
}
