package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMembersIncrementGetPrivy 增量获取卖家会员
// taobao.crm.members.increment.get.privy
//
// 增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100
func TaobaoCrmMembersIncrementGetPrivy(clt *core.SDKClient, req *crm.TaobaoCrmMembersIncrementGetPrivyAPIRequest, session string) (*crm.TaobaoCrmMembersIncrementGetPrivyAPIResponse, error) {
	var resp crm.TaobaoCrmMembersIncrementGetPrivyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
