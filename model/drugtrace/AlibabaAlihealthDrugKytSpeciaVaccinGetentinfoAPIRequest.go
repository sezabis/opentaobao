package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIRequest 通过企业名得到唯一标识(ref_ent_id)及企业ID(ent_id) API请求
// alibaba.alihealth.drug.kyt.specia.vaccin.getentinfo
//
// 根据企业名称查询企业唯一标识（ref_ent_id）及企业ID(ent_id)
type AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// NewAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoRequest 初始化AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytSpeciaVaccinGetentinfoRequest() *AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIRequest {
	return &AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.specia.vaccin.getentinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEntName is EntName Setter
// 公司名称(全称)
func (r *AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinGetentinfoAPIRequest) GetEntName() string {
	return r._entName
}
