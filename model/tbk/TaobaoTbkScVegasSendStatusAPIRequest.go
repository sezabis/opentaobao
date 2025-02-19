package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScVegasSendStatusAPIRequest 淘宝客-服务商-红包领取状态查询 API请求
// taobao.tbk.sc.vegas.send.status
//
// 服务商使用。支持淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包。接入前需签署协议：https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin
type TaobaoTbkScVegasSendStatusAPIRequest struct {
	model.Params
	// 会员运营id
	_specialId string
	// 渠道管理id
	_relationId string
	// 加密后的值(ALIPAY_ID除外)，需用MD5加密，32位小写
	_deviceValue string
	// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE 5. ALIPAY_ID
	_deviceType string
}

// NewTaobaoTbkScVegasSendStatusRequest 初始化TaobaoTbkScVegasSendStatusAPIRequest对象
func NewTaobaoTbkScVegasSendStatusRequest() *TaobaoTbkScVegasSendStatusAPIRequest {
	return &TaobaoTbkScVegasSendStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.vegas.send.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSpecialId is SpecialId Setter
// 会员运营id
func (r *TaobaoTbkScVegasSendStatusAPIRequest) SetSpecialId(_specialId string) error {
	r._specialId = _specialId
	r.Set("special_id", _specialId)
	return nil
}

// GetSpecialId SpecialId Getter
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetSpecialId() string {
	return r._specialId
}

// SetRelationId is RelationId Setter
// 渠道管理id
func (r *TaobaoTbkScVegasSendStatusAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetDeviceValue is DeviceValue Setter
// 加密后的值(ALIPAY_ID除外)，需用MD5加密，32位小写
func (r *TaobaoTbkScVegasSendStatusAPIRequest) SetDeviceValue(_deviceValue string) error {
	r._deviceValue = _deviceValue
	r.Set("device_value", _deviceValue)
	return nil
}

// GetDeviceValue DeviceValue Getter
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetDeviceValue() string {
	return r._deviceValue
}

// SetDeviceType is DeviceType Setter
// 入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE 5. ALIPAY_ID
func (r *TaobaoTbkScVegasSendStatusAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r TaobaoTbkScVegasSendStatusAPIRequest) GetDeviceType() string {
	return r._deviceType
}
