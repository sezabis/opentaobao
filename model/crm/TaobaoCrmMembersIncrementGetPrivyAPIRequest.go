package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMembersIncrementGetPrivyAPIRequest 增量获取卖家会员 API请求
// taobao.crm.members.increment.get.privy
//
// 增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100
type TaobaoCrmMembersIncrementGetPrivyAPIRequest struct {
	model.Params
	// 卖家修改会员信息的时间起点.
	_startModify string
	// 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间.
	_endModify string
	// 会员等级
	_grade int64
	// 每页显示的会员数，page_size的值不能超过100，最小值要大于1
	_pageSize int64
	// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1
	_currentPage int64
}

// NewTaobaoCrmMembersIncrementGetPrivyRequest 初始化TaobaoCrmMembersIncrementGetPrivyAPIRequest对象
func NewTaobaoCrmMembersIncrementGetPrivyRequest() *TaobaoCrmMembersIncrementGetPrivyAPIRequest {
	return &TaobaoCrmMembersIncrementGetPrivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMembersIncrementGetPrivyAPIRequest) GetApiMethodName() string {
	return "taobao.crm.members.increment.get.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMembersIncrementGetPrivyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStartModify is StartModify Setter
// 卖家修改会员信息的时间起点.
func (r *TaobaoCrmMembersIncrementGetPrivyAPIRequest) SetStartModify(_startModify string) error {
	r._startModify = _startModify
	r.Set("start_modify", _startModify)
	return nil
}

// GetStartModify StartModify Getter
func (r TaobaoCrmMembersIncrementGetPrivyAPIRequest) GetStartModify() string {
	return r._startModify
}

// SetEndModify is EndModify Setter
// 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间.
func (r *TaobaoCrmMembersIncrementGetPrivyAPIRequest) SetEndModify(_endModify string) error {
	r._endModify = _endModify
	r.Set("end_modify", _endModify)
	return nil
}

// GetEndModify EndModify Getter
func (r TaobaoCrmMembersIncrementGetPrivyAPIRequest) GetEndModify() string {
	return r._endModify
}

// SetGrade is Grade Setter
// 会员等级
func (r *TaobaoCrmMembersIncrementGetPrivyAPIRequest) SetGrade(_grade int64) error {
	r._grade = _grade
	r.Set("grade", _grade)
	return nil
}

// GetGrade Grade Getter
func (r TaobaoCrmMembersIncrementGetPrivyAPIRequest) GetGrade() int64 {
	return r._grade
}

// SetPageSize is PageSize Setter
// 每页显示的会员数，page_size的值不能超过100，最小值要大于1
func (r *TaobaoCrmMembersIncrementGetPrivyAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoCrmMembersIncrementGetPrivyAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1
func (r *TaobaoCrmMembersIncrementGetPrivyAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoCrmMembersIncrementGetPrivyAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
