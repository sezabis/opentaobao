package auction

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按年统计法院拍卖数据 APIResponse
taobao.auction.gov.data.annually.get

按月统计法院拍卖数据 包含：
标的件数统计：发布标的件数、结束标的件数、开拍标的件数
竞价实况：预计成交金额、出价次数、报名人数
在线标的：在线标的件数、意向用户数、网拍围观人次

最长6年，年起始时间2017年
*/
type TaobaoAuctionGovDataAnnuallyGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"auction_gov_data_annually_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 按年统计结果列表
    
    Results   []CourtsBidStatAnnuallyList `json:"results,omitempty" xml:"