package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthAdMediaDataSyncAPIRequest 媒体资源位投放效果数据回传 API请求
// taobao.usergrowth.ad.media.data.sync
//
// 创意维度广告效果数据回传
type TaobaoUsergrowthAdMediaDataSyncAPIRequest struct {
	model.Params
	// 统计日期
	_statDate string
	// 其他媒体
	_otherMedia string
	// 资源位
	_adPlacement string
	// 媒体
	_mediaCode string
	// 转化
	_transform int64
	// 消耗
	_cost int64
	// 曝光
	_exposure int64
	// 操作系统
	_deviceOs int64
	// 点击
	_click int64
	// 任务id
	_taskId int64
	// 渠道id
	_channelId int64
}

// NewTaobaoUsergrowthAdMediaDataSyncRequest 初始化TaobaoUsergrowthAdMediaDataSyncAPIRequest对象
func NewTaobaoUsergrowthAdMediaDataSyncRequest() *TaobaoUsergrowthAdMediaDataSyncAPIRequest {
	return &TaobaoUsergrowthAdMediaDataSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.ad.media.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStatDate is StatDate Setter
// 统计日期
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetStatDate(_statDate string) error {
	r._statDate = _statDate
	r.Set("stat_date", _statDate)
	return nil
}

// GetStatDate StatDate Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetStatDate() string {
	return r._statDate
}

// SetOtherMedia is OtherMedia Setter
// 其他媒体
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetOtherMedia(_otherMedia string) error {
	r._otherMedia = _otherMedia
	r.Set("other_media", _otherMedia)
	return nil
}

// GetOtherMedia OtherMedia Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetOtherMedia() string {
	return r._otherMedia
}

// SetAdPlacement is AdPlacement Setter
// 资源位
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetAdPlacement(_adPlacement string) error {
	r._adPlacement = _adPlacement
	r.Set("ad_placement", _adPlacement)
	return nil
}

// GetAdPlacement AdPlacement Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetAdPlacement() string {
	return r._adPlacement
}

// SetMediaCode is MediaCode Setter
// 媒体
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetMediaCode(_mediaCode string) error {
	r._mediaCode = _mediaCode
	r.Set("media_code", _mediaCode)
	return nil
}

// GetMediaCode MediaCode Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetMediaCode() string {
	return r._mediaCode
}

// SetTransform is Transform Setter
// 转化
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetTransform(_transform int64) error {
	r._transform = _transform
	r.Set("transform", _transform)
	return nil
}

// GetTransform Transform Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetTransform() int64 {
	return r._transform
}

// SetCost is Cost Setter
// 消耗
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetCost(_cost int64) error {
	r._cost = _cost
	r.Set("cost", _cost)
	return nil
}

// GetCost Cost Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetCost() int64 {
	return r._cost
}

// SetExposure is Exposure Setter
// 曝光
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetExposure(_exposure int64) error {
	r._exposure = _exposure
	r.Set("exposure", _exposure)
	return nil
}

// GetExposure Exposure Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetExposure() int64 {
	return r._exposure
}

// SetDeviceOs is DeviceOs Setter
// 操作系统
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetDeviceOs(_deviceOs int64) error {
	r._deviceOs = _deviceOs
	r.Set("device_os", _deviceOs)
	return nil
}

// GetDeviceOs DeviceOs Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetDeviceOs() int64 {
	return r._deviceOs
}

// SetClick is Click Setter
// 点击
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetClick(_click int64) error {
	r._click = _click
	r.Set("click", _click)
	return nil
}

// GetClick Click Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetClick() int64 {
	return r._click
}

// SetTaskId is TaskId Setter
// 任务id
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetTaskId() int64 {
	return r._taskId
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *TaobaoUsergrowthAdMediaDataSyncAPIRequest) SetChannelId(_channelId int64) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TaobaoUsergrowthAdMediaDataSyncAPIRequest) GetChannelId() int64 {
	return r._channelId
}
