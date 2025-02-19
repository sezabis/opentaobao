package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSearchbillDetailAPIRequest 查询单据详情 API请求
// alibaba.alihealth.drug.kyt.searchbill.detail
//
// 根据单据号码查询码单据详情和码信息
type AlibabaAlihealthDrugKytSearchbillDetailAPIRequest struct {
	model.Params
	// 单据号码
	_billCode string
	// 企业refEntId
	_refEntId string
}

// NewAlibabaAlihealthDrugKytSearchbillDetailRequest 初始化AlibabaAlihealthDrugKytSearchbillDetailAPIRequest对象
func NewAlibabaAlihealthDrugKytSearchbillDetailRequest() *AlibabaAlihealthDrugKytSearchbillDetailAPIRequest {
	return &AlibabaAlihealthDrugKytSearchbillDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSearchbillDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.searchbill.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSearchbillDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBillCode is BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugKytSearchbillDetailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytSearchbillDetailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetRefEntId is RefEntId Setter
// 企业refEntId
func (r *AlibabaAlihealthDrugKytSearchbillDetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytSearchbillDetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}
