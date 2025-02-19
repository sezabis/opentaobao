package alihealthmedical

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

// AlibabaAlihealthMedicalOrderRefuse 三方机构通知平台"医生拒诊"
// alibaba.alihealth.medical.order.refuse
//
// 三方机构通知平台&#34;医生拒诊&#34;
func AlibabaAlihealthMedicalOrderRefuse(clt *core.SDKClient, req *alihealthmedical.AlibabaAlihealthMedicalOrderRefuseAPIRequest, session string) (*alihealthmedical.AlibabaAlihealthMedicalOrderRefuseAPIResponse, error) {
	var resp alihealthmedical.AlibabaAlihealthMedicalOrderRefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
