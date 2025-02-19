package drugtrace

// LogCodeReplaceDto 结构体
type LogCodeReplaceDto struct {
	// 主键
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 操作人
	OperIcCode string `json:"oper_ic_code,omitempty" xml:"oper_ic_code,omitempty"`
	// 操作时间
	OperDate string `json:"oper_date,omitempty" xml:"oper_date,omitempty"`
	// 旧码
	PiatsCodeOld string `json:"piats_code_old,omitempty" xml:"piats_code_old,omitempty"`
	// 新码
	PiatsCodeNew string `json:"piats_code_new,omitempty" xml:"piats_code_new,omitempty"`
	// 企业ID
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 药品ID
	DrugEntBaseInfoId string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
	// 码级别
	CodeLevel int64 `json:"code_level,omitempty" xml:"code_level,omitempty"`
}
