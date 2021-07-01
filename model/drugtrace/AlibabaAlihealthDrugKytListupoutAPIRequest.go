package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytListupoutAPIRequest
查询货主/本企业上游企业出库单据信息 API请求
alibaba.alihealth.drug.kyt.listupout

查询货主/本企业上游企业出库单据信息 */
type AlibabaAlihealthDrugKytListupoutAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 开始日期（不写时分秒）
	_beginDate string
	// 结束日期（不写时分秒）
	_endDate string
	// 生产批号
	_produceBatchNo string
	// 药品ID
	_drugEntBaseInfoId string
	// 单据类型
	_billType string
	// 药品类型
	_physicType string
	// 状态
	_status string
	// 单据号
	_billCode string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
	// 发货单位
	_fromUserId string
}

// NewAlibabaAlihealthDrugKytListupoutRequest 初始化AlibabaAlihealthDrugKytListupoutAPIRequest对象
func NewAlibabaAlihealthDrugKytListupoutRequest() *AlibabaAlihealthDrugKytListupoutAPIRequest {
	return &AlibabaAlihealthDrugKytListupoutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.listupout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is BeginDate Setter
// 开始日期（不写时分秒）
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// Get BeginDate Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// Set is EndDate Setter
// 结束日期（不写时分秒）
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is ProduceBatchNo Setter
// 生产批号
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetProduceBatchNo(_produceBatchNo string) error {
	r._produceBatchNo = _produceBatchNo
	r.Set("produce_batch_no", _produceBatchNo)
	return nil
}

// Get ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetProduceBatchNo() string {
	return r._produceBatchNo
}

// Set is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// Get DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}

// Set is BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// Get BillType Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetBillType() string {
	return r._billType
}

// Set is PhysicType Setter
// 药品类型
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetPhysicType(_physicType string) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// Get PhysicType Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetPhysicType() string {
	return r._physicType
}

// Set is Status Setter
// 状态
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetStatus() string {
	return r._status
}

// Set is BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetPage() int64 {
	return r._page
}

// Set is FromUserId Setter
// 发货单位
func (r *AlibabaAlihealthDrugKytListupoutAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// Get FromUserId Getter
func (r AlibabaAlihealthDrugKytListupoutAPIRequest) GetFromUserId() string {
	return r._fromUserId
}
