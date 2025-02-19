package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangChannelInventoryBatchUpload ERP全量同步销售库存数量
// alibaba.dchain.aoxiang.channel.inventory.batch.upload
//
// ERP全量同步销售库存数量
func AlibabaDchainAoxiangChannelInventoryBatchUpload(clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangChannelInventoryBatchUploadAPIRequest, session string) (*ascp.AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse, error) {
	var resp ascp.AlibabaDchainAoxiangChannelInventoryBatchUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
