package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwGmAudit 同情用药审核接口
// alibaba.alihealth.pw.gm.audit
//
// 同情用药审核接口，提供给合作方审核申请单
func AlibabaAlihealthPwGmAudit(clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwGmAuditAPIRequest, session string) (*alihealthpw.AlibabaAlihealthPwGmAuditAPIResponse, error) {
	var resp alihealthpw.AlibabaAlihealthPwGmAuditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
