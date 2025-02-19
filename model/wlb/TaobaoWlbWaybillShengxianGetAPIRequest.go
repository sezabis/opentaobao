package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillShengxianGetAPIRequest 商家获取生鲜电子面单号 API请求
// taobao.wlb.waybill.shengxian.get
//
// 商家通过交易订单号获取电子面单接口
type TaobaoWlbWaybillShengxianGetAPIRequest struct {
	model.Params
	// 物流服务方代码，生鲜配送：YXSR
	_bizCode string
	// 物流服务类型。冷链：602，常温：502
	_deliveryType string
	// 商家淘宝地址信息ID
	_senderAddressId string
	// 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)
	_serviceCode string
	// 订单渠道： 淘宝平台订单："TB"; 天猫平台订单："TM"; 京东："JD"; 拍拍："PP"; 易迅平台订单："YX"; 一号店平台订单："YHD"; 当当网平台订单："DD"; EBAY："EBAY"; QQ网购："QQ"; 苏宁："SN"; 国美："GM"; 唯品会："WPH"; 聚美："Jm"; 乐峰："LF"; 蘑菇街："MGJ"; 聚尚："JS"; 银泰："YT"; VANCL："VANCL"; 邮乐："YL"; 优购："YG"; 拍鞋："PX"; 1688平台："1688";
	_orderChannelsType string
	// 交易号，传入交易号不能存在拆单场景。
	_tradeId string
	// 预留扩展字段
	_feature string
}

// NewTaobaoWlbWaybillShengxianGetRequest 初始化TaobaoWlbWaybillShengxianGetAPIRequest对象
func NewTaobaoWlbWaybillShengxianGetRequest() *TaobaoWlbWaybillShengxianGetAPIRequest {
	return &TaobaoWlbWaybillShengxianGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.shengxian.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizCode is BizCode Setter
// 物流服务方代码，生鲜配送：YXSR
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetDeliveryType is DeliveryType Setter
// 物流服务类型。冷链：602，常温：502
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetDeliveryType(_deliveryType string) error {
	r._deliveryType = _deliveryType
	r.Set("delivery_type", _deliveryType)
	return nil
}

// GetDeliveryType DeliveryType Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetDeliveryType() string {
	return r._deliveryType
}

// SetSenderAddressId is SenderAddressId Setter
// 商家淘宝地址信息ID
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetSenderAddressId(_senderAddressId string) error {
	r._senderAddressId = _senderAddressId
	r.Set("sender_address_id", _senderAddressId)
	return nil
}

// GetSenderAddressId SenderAddressId Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetSenderAddressId() string {
	return r._senderAddressId
}

// SetServiceCode is ServiceCode Setter
// 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

// SetOrderChannelsType is OrderChannelsType Setter
// 订单渠道： 淘宝平台订单：&#34;TB&#34;; 天猫平台订单：&#34;TM&#34;; 京东：&#34;JD&#34;; 拍拍：&#34;PP&#34;; 易迅平台订单：&#34;YX&#34;; 一号店平台订单：&#34;YHD&#34;; 当当网平台订单：&#34;DD&#34;; EBAY：&#34;EBAY&#34;; QQ网购：&#34;QQ&#34;; 苏宁：&#34;SN&#34;; 国美：&#34;GM&#34;; 唯品会：&#34;WPH&#34;; 聚美：&#34;Jm&#34;; 乐峰：&#34;LF&#34;; 蘑菇街：&#34;MGJ&#34;; 聚尚：&#34;JS&#34;; 银泰：&#34;YT&#34;; VANCL：&#34;VANCL&#34;; 邮乐：&#34;YL&#34;; 优购：&#34;YG&#34;; 拍鞋：&#34;PX&#34;; 1688平台：&#34;1688&#34;;
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetOrderChannelsType(_orderChannelsType string) error {
	r._orderChannelsType = _orderChannelsType
	r.Set("order_channels_type", _orderChannelsType)
	return nil
}

// GetOrderChannelsType OrderChannelsType Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetOrderChannelsType() string {
	return r._orderChannelsType
}

// SetTradeId is TradeId Setter
// 交易号，传入交易号不能存在拆单场景。
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetTradeId(_tradeId string) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetTradeId() string {
	return r._tradeId
}

// SetFeature is Feature Setter
// 预留扩展字段
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetFeature() string {
	return r._feature
}
