package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse 保险-开票结果反馈 API返回值
// taobao.servindustry.finance.insurance.invoice.feedback
//
// 保险-开票结果反馈
type TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponse struct {
	model.CommonResponse
	TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponseModel
}

// TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponseModel is 保险-开票结果反馈 成功返回结果
type TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"servindustry_finance_insurance_invoice_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult `json:"result,omitempty" xml:"result,omitempty"`
}
