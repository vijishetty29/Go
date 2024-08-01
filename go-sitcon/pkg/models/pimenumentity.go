package models

type PimEnumEntity struct {
	PK           int64  `gorm:"column:PK"`
	Code         string `gorm:"column:p_code"`
	ExternalCode string `gorm:"column:p_externalcode"`
}

func (PimEnumEntity) TableName() string {
	return "pimenumitem"
}

func GetPimEnumerationCodeByID(id int64) (string, error) {
	var enumeration *PimEnumEntity
	DB.Find(&enumeration, "PK = ?", id)

	if enumeration != nil {
		return enumeration.Code, nil
	}

	return "", nil
}
