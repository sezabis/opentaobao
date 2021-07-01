package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuTvDesktopToyouRecommendAPIRequest
TV桌面为你推荐接口 API请求
youku.tv.desktop.toyou.recommend

提供为你推荐数据 */
type YoukuTvDesktopToyouRecommendAPIRequest struct {
	model.Params
	// 用户登陆token
	_token string
	// 播控方，1:华数 7:CIBN
	_bcp string
	// 终端设备型号
	_deviceModel string
	// 桌面版本号
	_versionCode int64
	// 终端设备mac
	_mac string
	// 终端设备uuid
	_uuid string
	// 影视来源，0-淘tv 7-优酷免费 9-优酷付费 多项用“,”隔开
	_from string
	// 支持收费方式，0-免费 1-限免 2-单点 3-包月 4-红包 5-单包,多项用“,”隔开
	_chargeType string
	// 获取的最大节目数
	_maxSize int64
	// 分辨率，sw720 sw1080
	_sw string
	// 支持媒体类型,dts,drm,dolby,h265
	_deviceMedia string
	// 请求IP
	_ip string
}

// NewYoukuTvDesktopToyouRecommendRequest 初始化YoukuTvDesktopToyouRecommendAPIRequest对象
func NewYoukuTvDesktopToyouRecommendRequest() *YoukuTvDesktopToyouRecommendAPIRequest {
	return &YoukuTvDesktopToyouRecommendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetApiMethodName() string {
	return "youku.tv.desktop.toyou.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Token Setter
// 用户登陆token
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetToken() string {
	return r._token
}

// Set is Bcp Setter
// 播控方，1:华数 7:CIBN
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetBcp(_bcp string) error {
	r._bcp = _bcp
	r.Set("bcp", _bcp)
	return nil
}

// Get Bcp Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetBcp() string {
	return r._bcp
}

// Set is DeviceModel Setter
// 终端设备型号
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetDeviceModel(_deviceModel string) error {
	r._deviceModel = _deviceModel
	r.Set("device_model", _deviceModel)
	return nil
}

// Get DeviceModel Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetDeviceModel() string {
	return r._deviceModel
}

// Set is VersionCode Setter
// 桌面版本号
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetVersionCode(_versionCode int64) error {
	r._versionCode = _versionCode
	r.Set("version_code", _versionCode)
	return nil
}

// Get VersionCode Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetVersionCode() int64 {
	return r._versionCode
}

// Set is Mac Setter
// 终端设备mac
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// Get Mac Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetMac() string {
	return r._mac
}

// Set is Uuid Setter
// 终端设备uuid
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetUuid() string {
	return r._uuid
}

// Set is From Setter
// 影视来源，0-淘tv 7-优酷免费 9-优酷付费 多项用“,”隔开
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// Get From Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetFrom() string {
	return r._from
}

// Set is ChargeType Setter
// 支持收费方式，0-免费 1-限免 2-单点 3-包月 4-红包 5-单包,多项用“,”隔开
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetChargeType(_chargeType string) error {
	r._chargeType = _chargeType
	r.Set("charge_type", _chargeType)
	return nil
}

// Get ChargeType Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetChargeType() string {
	return r._chargeType
}

// Set is MaxSize Setter
// 获取的最大节目数
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetMaxSize(_maxSize int64) error {
	r._maxSize = _maxSize
	r.Set("max_size", _maxSize)
	return nil
}

// Get MaxSize Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetMaxSize() int64 {
	return r._maxSize
}

// Set is Sw Setter
// 分辨率，sw720 sw1080
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetSw(_sw string) error {
	r._sw = _sw
	r.Set("sw", _sw)
	return nil
}

// Get Sw Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetSw() string {
	return r._sw
}

// Set is DeviceMedia Setter
// 支持媒体类型,dts,drm,dolby,h265
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetDeviceMedia(_deviceMedia string) error {
	r._deviceMedia = _deviceMedia
	r.Set("device_media", _deviceMedia)
	return nil
}

// Get DeviceMedia Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetDeviceMedia() string {
	return r._deviceMedia
}

// Set is Ip Setter
// 请求IP
func (r *YoukuTvDesktopToyouRecommendAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// Get Ip Getter
func (r YoukuTvDesktopToyouRecommendAPIRequest) GetIp() string {
	return r._ip
}
