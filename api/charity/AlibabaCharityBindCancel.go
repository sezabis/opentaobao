package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityBindCancel 取消用户绑定
// alibaba.charity.bind.cancel
//
// 取消用户绑定
func AlibabaCharityBindCancel(clt *core.SDKClient, req *charity.AlibabaCharityBindCancelAPIRequest, session string) (*charity.AlibabaCharityBindCancelAPIResponse, error) {
	var resp charity.AlibabaCharityBindCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
