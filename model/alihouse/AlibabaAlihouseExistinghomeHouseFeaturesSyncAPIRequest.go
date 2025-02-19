package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest 房源标准打标数据同步 API请求
// alibaba.alihouse.existinghome.house.features.sync
//
// 房源标准打标数据同步
type AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest struct {
	model.Params
	// 房源上翻标
	_houseFeatures *SyncHouseFeaturesDto
}

// NewAlibabaAlihouseExistinghomeHouseFeaturesSyncRequest 初始化AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeHouseFeaturesSyncRequest() *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.features.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetHouseFeatures is HouseFeatures Setter
// 房源上翻标
func (r *AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) SetHouseFeatures(_houseFeatures *SyncHouseFeaturesDto) error {
	r._houseFeatures = _houseFeatures
	r.Set("house_features", _houseFeatures)
	return nil
}

// GetHouseFeatures HouseFeatures Getter
func (r AlibabaAlihouseExistinghomeHouseFeaturesSyncAPIRequest) GetHouseFeatures() *SyncHouseFeaturesDto {
	return r._houseFeatures
}
