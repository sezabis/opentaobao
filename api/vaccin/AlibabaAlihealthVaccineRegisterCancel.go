package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaAlihealthVaccineRegisterCancel 取消登记
// alibaba.alihealth.vaccine.register.cancel
//
// 取消登记
func AlibabaAlihealthVaccineRegisterCancel(clt *core.SDKClient, req *vaccin.AlibabaAlihealthVaccineRegisterCancelAPIRequest, session string) (*vaccin.AlibabaAlihealthVaccineRegisterCancelAPIResponse, error) {
	var resp vaccin.AlibabaAlihealthVaccineRegisterCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
