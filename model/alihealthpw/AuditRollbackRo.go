package alihealthpw

// AuditRollbackRo 结构体
type AuditRollbackRo struct {
	// 医院列表
	HospitalsDesc []string `json:"hospitals_desc,omitempty" xml:"hospitals_desc>string,omitempty"`
	// 打款日期
	ReceiptDate string `json:"receipt_date,omitempty" xml:"receipt_date,omitempty"`
	// 唯一编码
	UserUniqueCode string `json:"user_unique_code,omitempty" xml:"user_unique_code,omitempty"`
	// 审核时间
	ApplyAuditTime string `json:"apply_audit_time,omitempty" xml:"apply_audit_time,omitempty"`
	// xxx
	CheckRemark string `json:"check_remark,omitempty" xml:"check_remark,omitempty"`
	// 跳转url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 收款人
	ReceiptName string `json:"receipt_name,omitempty" xml:"receipt_name,omitempty"`
	// 收款账户
	ReceiptAccount string `json:"receipt_account,omitempty" xml:"receipt_account,omitempty"`
	// 审核状态
	ApplyAuditStatus string `json:"apply_audit_status,omitempty" xml:"apply_audit_status,omitempty"`
	// 申请项目
	ProjectThirdId string `json:"project_third_id,omitempty" xml:"project_third_id,omitempty"`
	// 收款金额
	ReceiptMoney string `json:"receipt_money,omitempty" xml:"receipt_money,omitempty"`
	// 患者姓名
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
	// 选择医院
	HospitalInfo *HospitalInfoDto `json:"hospital_info,omitempty" xml:"hospital_info,omitempty"`
}
