package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillIFullupdateAPIResponse 面单信息更新接口v1.0 API返回值
// taobao.wlb.waybill.i.fullupdate
//
// 商家更新电子面单号对应的订单信息。&lt;br/&gt;&lt;br/&gt;a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；&lt;br/&gt;&lt;br/&gt;b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。
type TaobaoWlbWaybillIFullupdateAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWaybillIFullupdateAPIResponseModel
}

// TaobaoWlbWaybillIFullupdateAPIResponseModel is 面单信息更新接口v1.0 成功返回结果
type TaobaoWlbWaybillIFullupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_waybill_i_fullupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新接口出参
	WaybillApplyUpdateInfo *WaybillApplyUpdateInfo `json:"waybill_apply_update_info,omitempty" xml:"waybill_apply_update_info,omitempty"`
}
