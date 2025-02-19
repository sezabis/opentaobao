package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest 根据码查询码信息（疫苗） API请求
// alibaba.alihealth.drug.code.kyt.specia.vaccin.querycode
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest struct {
	model.Params
	// 码列表
	_codes []string
	// 企业唯一标识（或appkey）
	_refEntId string
}

// NewAlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeRequest 初始化AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeRequest() *AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.specia.vaccin.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCodes is Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识（或appkey）
func (r *AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytSpeciaVaccinQuerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
