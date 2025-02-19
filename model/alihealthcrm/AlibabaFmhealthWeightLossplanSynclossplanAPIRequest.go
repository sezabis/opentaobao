package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFmhealthWeightLossplanSynclossplanAPIRequest 减重计划--同步减重计划 API请求
// alibaba.fmhealth.weight.lossplan.synclossplan
//
// 减重计划--三方同步用户初始化减重计划给我们
type AlibabaFmhealthWeightLossplanSynclossplanAPIRequest struct {
	model.Params
	// 生日
	_birthday string
	// 当前体重（今天的体重），单位kg
	_weight string
	// 目标体重
	_weightGoal string
	// 每周减重
	_lossPerWeek string
	// 完成时间
	_finishDate string
	// 减重计划开始时间
	_beginDate string
	// 用户id
	_tpUserId int64
	// 性别0女 1男
	_gender int64
	// 身高170 即一米七
	_height int64
	// 0创建减肥计划调用；1修改调用；
	_type int64
	// 减重类型0保持 1减肥
	_lossLevel int64
	// 体年龄
	_bodyAge int64
	// 每日可以摄入的标准总量
	_totalCalorie int64
}

// NewAlibabaFmhealthWeightLossplanSynclossplanRequest 初始化AlibabaFmhealthWeightLossplanSynclossplanAPIRequest对象
func NewAlibabaFmhealthWeightLossplanSynclossplanRequest() *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest {
	return &AlibabaFmhealthWeightLossplanSynclossplanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetApiMethodName() string {
	return "alibaba.fmhealth.weight.lossplan.synclossplan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBirthday is Birthday Setter
// 生日
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetBirthday(_birthday string) error {
	r._birthday = _birthday
	r.Set("birthday", _birthday)
	return nil
}

// GetBirthday Birthday Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetBirthday() string {
	return r._birthday
}

// SetWeight is Weight Setter
// 当前体重（今天的体重），单位kg
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetWeight(_weight string) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetWeight() string {
	return r._weight
}

// SetWeightGoal is WeightGoal Setter
// 目标体重
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetWeightGoal(_weightGoal string) error {
	r._weightGoal = _weightGoal
	r.Set("weight_goal", _weightGoal)
	return nil
}

// GetWeightGoal WeightGoal Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetWeightGoal() string {
	return r._weightGoal
}

// SetLossPerWeek is LossPerWeek Setter
// 每周减重
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetLossPerWeek(_lossPerWeek string) error {
	r._lossPerWeek = _lossPerWeek
	r.Set("loss_per_week", _lossPerWeek)
	return nil
}

// GetLossPerWeek LossPerWeek Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetLossPerWeek() string {
	return r._lossPerWeek
}

// SetFinishDate is FinishDate Setter
// 完成时间
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetFinishDate(_finishDate string) error {
	r._finishDate = _finishDate
	r.Set("finish_date", _finishDate)
	return nil
}

// GetFinishDate FinishDate Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetFinishDate() string {
	return r._finishDate
}

// SetBeginDate is BeginDate Setter
// 减重计划开始时间
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetTpUserId is TpUserId Setter
// 用户id
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetTpUserId(_tpUserId int64) error {
	r._tpUserId = _tpUserId
	r.Set("tp_user_id", _tpUserId)
	return nil
}

// GetTpUserId TpUserId Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetTpUserId() int64 {
	return r._tpUserId
}

// SetGender is Gender Setter
// 性别0女 1男
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetGender(_gender int64) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// GetGender Gender Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetGender() int64 {
	return r._gender
}

// SetHeight is Height Setter
// 身高170 即一米七
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetHeight(_height int64) error {
	r._height = _height
	r.Set("height", _height)
	return nil
}

// GetHeight Height Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetHeight() int64 {
	return r._height
}

// SetType is Type Setter
// 0创建减肥计划调用；1修改调用；
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetType() int64 {
	return r._type
}

// SetLossLevel is LossLevel Setter
// 减重类型0保持 1减肥
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetLossLevel(_lossLevel int64) error {
	r._lossLevel = _lossLevel
	r.Set("loss_level", _lossLevel)
	return nil
}

// GetLossLevel LossLevel Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetLossLevel() int64 {
	return r._lossLevel
}

// SetBodyAge is BodyAge Setter
// 体年龄
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetBodyAge(_bodyAge int64) error {
	r._bodyAge = _bodyAge
	r.Set("body_age", _bodyAge)
	return nil
}

// GetBodyAge BodyAge Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetBodyAge() int64 {
	return r._bodyAge
}

// SetTotalCalorie is TotalCalorie Setter
// 每日可以摄入的标准总量
func (r *AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) SetTotalCalorie(_totalCalorie int64) error {
	r._totalCalorie = _totalCalorie
	r.Set("total_calorie", _totalCalorie)
	return nil
}

// GetTotalCalorie TotalCalorie Getter
func (r AlibabaFmhealthWeightLossplanSynclossplanAPIRequest) GetTotalCalorie() int64 {
	return r._totalCalorie
}
