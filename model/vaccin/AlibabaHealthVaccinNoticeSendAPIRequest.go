package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeSendAPIRequest 发送消息提醒 API请求
// alibaba.health.vaccin.notice.send
//
// ISV 通过免疫规划中心给用户发送短信或者支付宝 PUSH 提醒。
type AlibabaHealthVaccinNoticeSendAPIRequest struct {
	model.Params
	// 接种的疫苗信息
	_vaccineList []VaccineInfo
	// 接种人出生日期
	_birthday string
	// 区名称
	_area string
	// 接种点地址
	_address string
	// 订单 ID
	_orderId string
	// 接种点名称
	_povName string
	// 城市名称
	_city string
	// 接种点编码
	_povNo string
	// 联系电话
	_mobile string
	// 省份名称
	_province string
	// 预约时间段
	_reserveTime string
	// 接种人姓名
	_name string
	// ISV 侧用户 ID
	_isvUserId string
	// 预约日期
	_reserveDate string
	// 支付宝用户 ID
	_alipayUserId string
	// 用户入口，支付宝或医鹿，alipay或yl
	_appChannel string
	// 接种人性别:1=男,2=女
	_sex int64
	// 消息提醒类型：1=疫苗下一针预约提醒  2=预约成功提醒  3=接种提醒  4=补种提醒
	_messageType int64
}

// NewAlibabaHealthVaccinNoticeSendRequest 初始化AlibabaHealthVaccinNoticeSendAPIRequest对象
func NewAlibabaHealthVaccinNoticeSendRequest() *AlibabaHealthVaccinNoticeSendAPIRequest {
	return &AlibabaHealthVaccinNoticeSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetVaccineList is VaccineList Setter
// 接种的疫苗信息
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetVaccineList(_vaccineList []VaccineInfo) error {
	r._vaccineList = _vaccineList
	r.Set("vaccine_list", _vaccineList)
	return nil
}

// GetVaccineList VaccineList Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetVaccineList() []VaccineInfo {
	return r._vaccineList
}

// SetBirthday is Birthday Setter
// 接种人出生日期
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetBirthday(_birthday string) error {
	r._birthday = _birthday
	r.Set("birthday", _birthday)
	return nil
}

// GetBirthday Birthday Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetBirthday() string {
	return r._birthday
}

// SetArea is Area Setter
// 区名称
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// GetArea Area Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetArea() string {
	return r._area
}

// SetAddress is Address Setter
// 接种点地址
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetAddress() string {
	return r._address
}

// SetOrderId is OrderId Setter
// 订单 ID
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetPovName is PovName Setter
// 接种点名称
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetPovName(_povName string) error {
	r._povName = _povName
	r.Set("pov_name", _povName)
	return nil
}

// GetPovName PovName Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetPovName() string {
	return r._povName
}

// SetCity is City Setter
// 城市名称
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetCity() string {
	return r._city
}

// SetPovNo is PovNo Setter
// 接种点编码
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetPovNo(_povNo string) error {
	r._povNo = _povNo
	r.Set("pov_no", _povNo)
	return nil
}

// GetPovNo PovNo Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetPovNo() string {
	return r._povNo
}

// SetMobile is Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetMobile() string {
	return r._mobile
}

// SetProvince is Province Setter
// 省份名称
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetProvince() string {
	return r._province
}

// SetReserveTime is ReserveTime Setter
// 预约时间段
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetReserveTime(_reserveTime string) error {
	r._reserveTime = _reserveTime
	r.Set("reserve_time", _reserveTime)
	return nil
}

// GetReserveTime ReserveTime Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetReserveTime() string {
	return r._reserveTime
}

// SetName is Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetName() string {
	return r._name
}

// SetIsvUserId is IsvUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetIsvUserId(_isvUserId string) error {
	r._isvUserId = _isvUserId
	r.Set("isv_user_id", _isvUserId)
	return nil
}

// GetIsvUserId IsvUserId Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetIsvUserId() string {
	return r._isvUserId
}

// SetReserveDate is ReserveDate Setter
// 预约日期
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetReserveDate(_reserveDate string) error {
	r._reserveDate = _reserveDate
	r.Set("reserve_date", _reserveDate)
	return nil
}

// GetReserveDate ReserveDate Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetReserveDate() string {
	return r._reserveDate
}

// SetAlipayUserId is AlipayUserId Setter
// 支付宝用户 ID
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// SetAppChannel is AppChannel Setter
// 用户入口，支付宝或医鹿，alipay或yl
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetAppChannel(_appChannel string) error {
	r._appChannel = _appChannel
	r.Set("app_channel", _appChannel)
	return nil
}

// GetAppChannel AppChannel Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetAppChannel() string {
	return r._appChannel
}

// SetSex is Sex Setter
// 接种人性别:1=男,2=女
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetSex(_sex int64) error {
	r._sex = _sex
	r.Set("sex", _sex)
	return nil
}

// GetSex Sex Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetSex() int64 {
	return r._sex
}

// SetMessageType is MessageType Setter
// 消息提醒类型：1=疫苗下一针预约提醒  2=预约成功提醒  3=接种提醒  4=补种提醒
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetMessageType(_messageType int64) error {
	r._messageType = _messageType
	r.Set("message_type", _messageType)
	return nil
}

// GetMessageType MessageType Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetMessageType() int64 {
	return r._messageType
}
