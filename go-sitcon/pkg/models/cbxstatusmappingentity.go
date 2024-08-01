package models

type CbxStatusMappingEntity struct {
	PK                 int64 `gorm:"column:pk"`
	WorkflowStatus     int64 `gorm:"column:p_workflowstatus"`
	DocumentStatus     int64 `gorm:"column:p_documentstatus"`
	CountryOrderStatus int64 `gorm:"column:p_countryorderstatus"`
	PSAStatus          int64 `gorm:"column:p_psastatus"`
}

func (CbxStatusMappingEntity) TableName() string {
	return "PimCbxStatusComboMapping"
}

func GetPSAStatusForWorkflowDocumentAndCountryOrderStatus(workflowStatus int64, documentStatus int64, countryOrderStatus int64) int64 {
	var cbxStatus *CbxStatusMappingEntity
	DB.Find(&cbxStatus, "p_workflowstatus = ? AND p_documentstatus = ? AND p_countryorderstatus = ?", workflowStatus, documentStatus, countryOrderStatus)
	return cbxStatus.PSAStatus
}
