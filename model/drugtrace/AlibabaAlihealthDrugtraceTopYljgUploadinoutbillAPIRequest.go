package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest
出入库单据上传 API请求
alibaba.alihealth.drugtrace.top.yljg.uploadinoutbill

零售企业上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）,
不包括对个人的零售出库，疫苗接种，领药出库。 */
type AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest struct {
	model.Params
	// 单据编码
	_billCode string
	// 单据时间
	_billTime string
	// 单据类型【102代表采购入库】
	_billType int64
	// 药品类型【3普药2特药】
	_physicType int64
	// 上传企业的单位编码
	_refUserId string
	// 代理企业REF标识
	_agentRefUserId string
	// 发货企业entId
	_fromUserId string
	// 收货企业entId
	_toUserId string
	// 直调企业标识
	_destUserId string
	// 单据提交者（appkey编号）
	_operIcCode string
	// 单据提交者姓名
	_operIcName string
	// 客户端类型[必须填2]
	_clientType string
	// 退货原因代码[退货入出库时填写]
	_returnReasonCode string
	// 退货原因描述[退货入出库时填写]
	_returnReasonDes string
	// 注销原因代码【销毁出库时填写】
	_cancelReasonCode string
	// 注销原因描述【销毁出库时填写】
	_cancelReasonDes string
	// 执行人姓名【销毁出库时填写】
	_executerName string
	// 执行人证件号【销毁出库时填写】
	_executerCode string
	// 监督人姓名【销毁出库时填写】
	_superviserName string
	// 监督人证件号【销毁出库时填写】
	_superviserCode string
	// 仓号
	_warehouseId string
	// 药品ID[企业自已系统的药品ID]
	_drugId string
	// 追溯码[多个时用逗号分开]
	_traceCodes []string
	// （协同平台数据合规）发货地址【必选】
	_fromAddress string
	// （协同平台数据合规）收货地址【必选】
	_toAddress string
	// （协同平台数据合规）发货单编号【必选】
	_fromBillCode string
	// （协同平台数据合规）订货单编号【可选】
	_orderCode string
	// （协同平台数据合规）发货人【必选】
	_fromPerson string
	// （协同平台数据合规）收货人【必选】
	_toPerson string
	// （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
	_disRefEntId string
	// （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
	_disEntId string
	// （协同平台数据合规）应收货总数量【必选】
	_quReceivable int64
	// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
	_xtIsCheck string
	// （协同平台数据合规）未验证通过原因【验证未通过时填写】
	_xtCheckCode string
	// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
	_xtCheckCodeDesc string
	// （协同平台数据合规）药品列表Json
	_drugListJson string
	// （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】
	_assRefEntId string
	// （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】
	_assEntId string
}

// NewAlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest 初始化AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgUploadinoutbillRequest() *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.uploadinoutbill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is BillTime Setter
// 单据时间
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// Get BillTime Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetBillTime() string {
	return r._billTime
}

// Set is BillType Setter
// 单据类型【102代表采购入库】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// Get BillType Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetBillType() int64 {
	return r._billType
}

// Set is PhysicType Setter
// 药品类型【3普药2特药】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// Get PhysicType Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

// Set is RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// Get RefUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// Set is AgentRefUserId Setter
// 代理企业REF标识
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetAgentRefUserId(_agentRefUserId string) error {
	r._agentRefUserId = _agentRefUserId
	r.Set("agent_ref_user_id", _agentRefUserId)
	return nil
}

// Get AgentRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetAgentRefUserId() string {
	return r._agentRefUserId
}

// Set is FromUserId Setter
// 发货企业entId
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// Get FromUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// Set is ToUserId Setter
// 收货企业entId
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetToUserId(_toUserId string) error {
	r._toUserId = _toUserId
	r.Set("to_user_id", _toUserId)
	return nil
}

// Get ToUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetToUserId() string {
	return r._toUserId
}

// Set is DestUserId Setter
// 直调企业标识
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetDestUserId(_destUserId string) error {
	r._destUserId = _destUserId
	r.Set("dest_user_id", _destUserId)
	return nil
}

// Get DestUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetDestUserId() string {
	return r._destUserId
}

// Set is OperIcCode Setter
// 单据提交者（appkey编号）
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// Get OperIcCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// Set is OperIcName Setter
// 单据提交者姓名
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// Get OperIcName Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// Set is ClientType Setter
// 客户端类型[必须填2]
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// Get ClientType Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetClientType() string {
	return r._clientType
}

// Set is ReturnReasonCode Setter
// 退货原因代码[退货入出库时填写]
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetReturnReasonCode(_returnReasonCode string) error {
	r._returnReasonCode = _returnReasonCode
	r.Set("return_reason_code", _returnReasonCode)
	return nil
}

// Get ReturnReasonCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetReturnReasonCode() string {
	return r._returnReasonCode
}

// Set is ReturnReasonDes Setter
// 退货原因描述[退货入出库时填写]
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetReturnReasonDes(_returnReasonDes string) error {
	r._returnReasonDes = _returnReasonDes
	r.Set("return_reason_des", _returnReasonDes)
	return nil
}

// Get ReturnReasonDes Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetReturnReasonDes() string {
	return r._returnReasonDes
}

// Set is CancelReasonCode Setter
// 注销原因代码【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetCancelReasonCode(_cancelReasonCode string) error {
	r._cancelReasonCode = _cancelReasonCode
	r.Set("cancel_reason_code", _cancelReasonCode)
	return nil
}

// Get CancelReasonCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetCancelReasonCode() string {
	return r._cancelReasonCode
}

// Set is CancelReasonDes Setter
// 注销原因描述【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetCancelReasonDes(_cancelReasonDes string) error {
	r._cancelReasonDes = _cancelReasonDes
	r.Set("cancel_reason_des", _cancelReasonDes)
	return nil
}

// Get CancelReasonDes Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetCancelReasonDes() string {
	return r._cancelReasonDes
}

// Set is ExecuterName Setter
// 执行人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetExecuterName(_executerName string) error {
	r._executerName = _executerName
	r.Set("executer_name", _executerName)
	return nil
}

// Get ExecuterName Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetExecuterName() string {
	return r._executerName
}

// Set is ExecuterCode Setter
// 执行人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetExecuterCode(_executerCode string) error {
	r._executerCode = _executerCode
	r.Set("executer_code", _executerCode)
	return nil
}

// Get ExecuterCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetExecuterCode() string {
	return r._executerCode
}

// Set is SuperviserName Setter
// 监督人姓名【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetSuperviserName(_superviserName string) error {
	r._superviserName = _superviserName
	r.Set("superviser_name", _superviserName)
	return nil
}

// Get SuperviserName Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetSuperviserName() string {
	return r._superviserName
}

// Set is SuperviserCode Setter
// 监督人证件号【销毁出库时填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetSuperviserCode(_superviserCode string) error {
	r._superviserCode = _superviserCode
	r.Set("superviser_code", _superviserCode)
	return nil
}

// Get SuperviserCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetSuperviserCode() string {
	return r._superviserCode
}

// Set is WarehouseId Setter
// 仓号
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetWarehouseId(_warehouseId string) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// Get WarehouseId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetWarehouseId() string {
	return r._warehouseId
}

// Set is DrugId Setter
// 药品ID[企业自已系统的药品ID]
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetDrugId(_drugId string) error {
	r._drugId = _drugId
	r.Set("drug_id", _drugId)
	return nil
}

// Get DrugId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetDrugId() string {
	return r._drugId
}

// Set is TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetTraceCodes(_traceCodes []string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// Get TraceCodes Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetTraceCodes() []string {
	return r._traceCodes
}

// Set is FromAddress Setter
// （协同平台数据合规）发货地址【必选】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetFromAddress(_fromAddress string) error {
	r._fromAddress = _fromAddress
	r.Set("from_address", _fromAddress)
	return nil
}

// Get FromAddress Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetFromAddress() string {
	return r._fromAddress
}

// Set is ToAddress Setter
// （协同平台数据合规）收货地址【必选】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetToAddress(_toAddress string) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// Get ToAddress Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetToAddress() string {
	return r._toAddress
}

// Set is FromBillCode Setter
// （协同平台数据合规）发货单编号【必选】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetFromBillCode(_fromBillCode string) error {
	r._fromBillCode = _fromBillCode
	r.Set("from_bill_code", _fromBillCode)
	return nil
}

// Get FromBillCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetFromBillCode() string {
	return r._fromBillCode
}

// Set is OrderCode Setter
// （协同平台数据合规）订货单编号【可选】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// Get OrderCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// Set is FromPerson Setter
// （协同平台数据合规）发货人【必选】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetFromPerson(_fromPerson string) error {
	r._fromPerson = _fromPerson
	r.Set("from_person", _fromPerson)
	return nil
}

// Get FromPerson Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetFromPerson() string {
	return r._fromPerson
}

// Set is ToPerson Setter
// （协同平台数据合规）收货人【必选】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetToPerson(_toPerson string) error {
	r._toPerson = _toPerson
	r.Set("to_person", _toPerson)
	return nil
}

// Get ToPerson Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetToPerson() string {
	return r._toPerson
}

// Set is DisRefEntId Setter
// （协同平台数据合规）药品配送企业【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetDisRefEntId(_disRefEntId string) error {
	r._disRefEntId = _disRefEntId
	r.Set("dis_ref_ent_id", _disRefEntId)
	return nil
}

// Get DisRefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetDisRefEntId() string {
	return r._disRefEntId
}

// Set is DisEntId Setter
// （协同平台数据合规）药品配送企业entId【出库单，收货方为医疗机构时填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetDisEntId(_disEntId string) error {
	r._disEntId = _disEntId
	r.Set("dis_ent_id", _disEntId)
	return nil
}

// Get DisEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetDisEntId() string {
	return r._disEntId
}

// Set is QuReceivable Setter
// （协同平台数据合规）应收货总数量【必选】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetQuReceivable(_quReceivable int64) error {
	r._quReceivable = _quReceivable
	r.Set("qu_receivable", _quReceivable)
	return nil
}

// Get QuReceivable Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetQuReceivable() int64 {
	return r._quReceivable
}

// Set is XtIsCheck Setter
// （协同平台数据合规）是否验证，0：未通过验证，1：已验证
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetXtIsCheck(_xtIsCheck string) error {
	r._xtIsCheck = _xtIsCheck
	r.Set("xt_is_check", _xtIsCheck)
	return nil
}

// Get XtIsCheck Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetXtIsCheck() string {
	return r._xtIsCheck
}

// Set is XtCheckCode Setter
// （协同平台数据合规）未验证通过原因【验证未通过时填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetXtCheckCode(_xtCheckCode string) error {
	r._xtCheckCode = _xtCheckCode
	r.Set("xt_check_code", _xtCheckCode)
	return nil
}

// Get XtCheckCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetXtCheckCode() string {
	return r._xtCheckCode
}

// Set is XtCheckCodeDesc Setter
// （协同平台数据合规）未验证通过原因描述【验证未通过时填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetXtCheckCodeDesc(_xtCheckCodeDesc string) error {
	r._xtCheckCodeDesc = _xtCheckCodeDesc
	r.Set("xt_check_code_desc", _xtCheckCodeDesc)
	return nil
}

// Get XtCheckCodeDesc Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetXtCheckCodeDesc() string {
	return r._xtCheckCodeDesc
}

// Set is DrugListJson Setter
// （协同平台数据合规）药品列表Json
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetDrugListJson(_drugListJson string) error {
	r._drugListJson = _drugListJson
	r.Set("drug_list_json", _drugListJson)
	return nil
}

// Get DrugListJson Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetDrugListJson() string {
	return r._drugListJson
}

// Set is AssRefEntId Setter
// （协同平台数据合规）单据委托企业refEntId【疫苗药品出库单填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetAssRefEntId(_assRefEntId string) error {
	r._assRefEntId = _assRefEntId
	r.Set("ass_ref_ent_id", _assRefEntId)
	return nil
}

// Get AssRefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetAssRefEntId() string {
	return r._assRefEntId
}

// Set is AssEntId Setter
// （协同平台数据合规）单据委托企业entId【疫苗药品出库单填写】
func (r *AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) SetAssEntId(_assEntId string) error {
	r._assEntId = _assEntId
	r.Set("ass_ent_id", _assEntId)
	return nil
}

// Get AssEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIRequest) GetAssEntId() string {
	return r._assEntId
}
