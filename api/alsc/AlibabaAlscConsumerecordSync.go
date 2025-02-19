package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscConsumerecordSync 外域订单同步本地生活消费记录
// alibaba.alsc.consumerecord.sync
//
// 外部第三方将本地生活app端下单数据同步到本地生活消费记录中心
func AlibabaAlscConsumerecordSync(clt *core.SDKClient, req *alsc.AlibabaAlscConsumerecordSyncAPIRequest, session string) (*alsc.AlibabaAlscConsumerecordSyncAPIResponse, error) {
	var resp alsc.AlibabaAlscConsumerecordSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
