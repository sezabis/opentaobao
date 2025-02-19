package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrDonateInvoiceQuerytoblockchainoss 触发odps任务离线查询公益宝贝开票对账明细
// alibaba.csr.donate.invoice.querytoblockchainoss
//
// 提供给蚂蚁链上公益团队，用于触发odps任务离线查询公益宝贝开票对账明细
func AlibabaCsrDonateInvoiceQuerytoblockchainoss(clt *core.SDKClient, req *charity.AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest, session string) (*charity.AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse, error) {
	var resp charity.AlibabaCsrDonateInvoiceQuerytoblockchainossAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
