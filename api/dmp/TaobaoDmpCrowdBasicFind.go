package dmp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dmp"
)

// TaobaoDmpCrowdBasicFind DMP_BP版人群列表查询
// taobao.dmp.crowd.basic.find
//
// DMP_BP版人群列表查询
func TaobaoDmpCrowdBasicFind(clt *core.SDKClient, req *dmp.TaobaoDmpCrowdBasicFindAPIRequest, session string) (*dmp.TaobaoDmpCrowdBasicFindAPIResponse, error) {
	var resp dmp.TaobaoDmpCrowdBasicFindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
