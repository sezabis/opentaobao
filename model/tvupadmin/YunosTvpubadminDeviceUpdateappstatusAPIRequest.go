package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceUpdateappstatusAPIRequest
更新应用版本审核状态 API请求
yunos.tvpubadmin.device.updateappstatus

更新应用版本审核状态 */
type YunosTvpubadminDeviceUpdateappstatusAPIRequest struct {
	model.Params
	// 应用ID
	_versionId int64
	// 牌照方
	_license int64
	// 审核状态
	_status string
	// 审核意见
	_auditComment string
}

// NewYunosTvpubadminDeviceUpdateappstatusRequest 初始化YunosTvpubadminDeviceUpdateappstatusAPIRequest对象
func NewYunosTvpubadminDeviceUpdateappstatusRequest() *YunosTvpubadminDeviceUpdateappstatusAPIRequest {
	return &YunosTvpubadminDeviceUpdateappstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceUpdateappstatusAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.updateappstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceUpdateappstatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is VersionId Setter
// 应用ID
func (r *YunosTvpubadminDeviceUpdateappstatusAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// Get VersionId Getter
func (r YunosTvpubadminDeviceUpdateappstatusAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// Set is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceUpdateappstatusAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// Get License Getter
func (r YunosTvpubadminDeviceUpdateappstatusAPIRequest) GetLicense() int64 {
	return r._license
}

// Set is Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceUpdateappstatusAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r YunosTvpubadminDeviceUpdateappstatusAPIRequest) GetStatus() string {
	return r._status
}

// Set is AuditComment Setter
// 审核意见
func (r *YunosTvpubadminDeviceUpdateappstatusAPIRequest) SetAuditComment(_auditComment string) error {
	r._auditComment = _auditComment
	r.Set("audit_comment", _auditComment)
	return nil
}

// Get AuditComment Getter
func (r YunosTvpubadminDeviceUpdateappstatusAPIRequest) GetAuditComment() string {
	return r._auditComment
}
