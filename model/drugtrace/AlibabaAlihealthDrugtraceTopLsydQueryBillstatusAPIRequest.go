package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest
上传单据后处理状态查询 API请求
alibaba.alihealth.drugtrace.top.lsyd.query.billstatus

单据处理状态查询 */
type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 开始日期
	_beginDate string
	// 结束日期
	_endDate string
	// 单据类型 A：全部 AI：全部入库 AO：全部出库
	_billType string
	// 单据号
	_billCode string
	// 药品类型
	_drugType string
	// 状态  0, 上传成功     3, 处理成功     4, 处理失败
	_dealStatus string
	// 发货商
	_fromUserId string
	// 收货商
	_toUserId string
	// 代理商
	_agentRefUserId string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
}

// NewAlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest 初始化AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryBillstatusRequest() *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest {
	return &AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.billstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is BeginDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// Get BeginDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// Set is EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is BillType Setter
// 单据类型 A：全部 AI：全部入库 AO：全部出库
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// Get BillType Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetBillType() string {
	return r._billType
}

// Set is BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is DrugType Setter
// 药品类型
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetDrugType(_drugType string) error {
	r._drugType = _drugType
	r.Set("drug_type", _drugType)
	return nil
}

// Get DrugType Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetDrugType() string {
	return r._drugType
}

// Set is DealStatus Setter
// 状态  0, 上传成功     3, 处理成功     4, 处理失败
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetDealStatus(_dealStatus string) error {
	r._dealStatus = _dealStatus
	r.Set("deal_status", _dealStatus)
	return nil
}

// Get DealStatus Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetDealStatus() string {
	return r._dealStatus
}

// Set is FromUserId Setter
// 发货商
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// Get FromUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// Set is ToUserId Setter
// 收货商
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// Get ToUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetToUserId() string {
	return r._toUserId
}

// Set is AgentRefUserId Setter
// 代理商
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// Get AgentRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// Set is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is Page Setter
// 页码
func (r *AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest) GetPage() int64 {
	return r._page
}
