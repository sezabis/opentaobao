package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest 查询追溯码对应的药品信息 API请求
// alibaba.alihealth.drug.code.kyt.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest struct {
	model.Params
	// 码列表【多个码用逗号分隔的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_codes []string
	// 企业唯一标识
	_refEntId string
}

// NewAlibabaAlihealthDrugCodeKytQuerycodeRequest 初始化AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytQuerycodeRequest() *AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCodes is Codes Setter
// 码列表【多个码用逗号分隔的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}
