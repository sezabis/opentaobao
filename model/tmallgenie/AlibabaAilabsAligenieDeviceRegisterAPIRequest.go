package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieDeviceRegisterAPIRequest
天猫精灵开放平台获取设备秘钥 API请求
alibaba.ailabs.aligenie.device.register

向天猫精灵inside平台注册设备mac地址，并获取设备的唯一密钥 */
type AlibabaAilabsAligenieDeviceRegisterAPIRequest struct {
	model.Params
	// 设备id
	_deviceId int64
	// mac区段脚本
	_macSections string
}

// NewAlibabaAilabsAligenieDeviceRegisterRequest 初始化AlibabaAilabsAligenieDeviceRegisterAPIRequest对象
func NewAlibabaAilabsAligenieDeviceRegisterRequest() *AlibabaAilabsAligenieDeviceRegisterAPIRequest {
	return &AlibabaAilabsAligenieDeviceRegisterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieDeviceRegisterAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.device.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieDeviceRegisterAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeviceId Setter
// 设备id
func (r *AlibabaAilabsAligenieDeviceRegisterAPIRequest) SetDeviceId(_deviceId int64) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// Get DeviceId Getter
func (r AlibabaAilabsAligenieDeviceRegisterAPIRequest) GetDeviceId() int64 {
	return r._deviceId
}

// Set is MacSections Setter
// mac区段脚本
func (r *AlibabaAilabsAligenieDeviceRegisterAPIRequest) SetMacSections(_macSections string) error {
	r._macSections = _macSections
	r.Set("mac_sections", _macSections)
	return nil
}

// Get MacSections Getter
func (r AlibabaAilabsAligenieDeviceRegisterAPIRequest) GetMacSections() string {
	return r._macSections
}
