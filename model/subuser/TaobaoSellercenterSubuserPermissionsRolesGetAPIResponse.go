package subuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse 查询指定的子账号的权限和角色信息 API返回值
// taobao.sellercenter.subuser.permissions.roles.get
//
// 查询指定的子账号的被直接赋予的权限信息和角色信息。&lt;br/&gt;返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
type TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse struct {
	model.CommonResponse
	TaobaoSellercenterSubuserPermissionsRolesGetAPIResponseModel
}

// TaobaoSellercenterSubuserPermissionsRolesGetAPIResponseModel is 查询指定的子账号的权限和角色信息 成功返回结果
type TaobaoSellercenterSubuserPermissionsRolesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"sellercenter_subuser_permissions_roles_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号被所拥有的权限
	SubuserPermission *SubUserPermission `json:"subuser_permission,omitempty" xml:"subuser_permission,omitempty"`
}
