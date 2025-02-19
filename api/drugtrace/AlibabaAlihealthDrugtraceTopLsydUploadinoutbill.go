package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydUploadinoutbill 出入库单据上传
// alibaba.alihealth.drugtrace.top.lsyd.uploadinoutbill
//
// 零售企业上传出入库信息，包括102, &#34;采购入库&#34;；103, &#34;退货入库&#34;；104, &#34;调拨入库&#34;；107, &#34;供应入库&#34;；108, &#34;召回入库&#34;；110,&#34;赠品入库&#34;；111,&#34;盘盈入库&#34;；112,&#34;报废入库&#34;；113,&#34;其他入库&#34;
// 201, &#34;销售出库&#34;；202, &#34;退货出库&#34;；203, &#34;调拨出库&#34;；204, &#34;返工出库&#34;；205, &#34;销毁出库&#34;；206, &#34;抽检出库&#34;；207, &#34;直调出库&#34;；209, &#34;供应出库&#34;；211, &#34;召回出库&#34;；212,&#34;赠品出库&#34;；214,&#34;盘亏出库&#34;；215,&#34;损坏出库&#34;；216,&#34;报废出库&#34;；217,&#34;其他出库&#34;；237, &#34;直调退货&#34;。
// 不包括对个人的零售出库，疫苗接种，领药出库。
func AlibabaAlihealthDrugtraceTopLsydUploadinoutbill(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
