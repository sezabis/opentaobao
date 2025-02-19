package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerRegister 服务商添加工人
// alibaba.ssc.supplyplatform.serviceworker.register
//
// 工人注册
func AlibabaSscSupplyplatformServiceworkerRegister(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerRegisterAPIRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse, error) {
	var resp tmallservice.AlibabaSscSupplyplatformServiceworkerRegisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
