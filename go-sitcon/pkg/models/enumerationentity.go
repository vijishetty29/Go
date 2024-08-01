package models

type EnumerationEntity struct {
	ID            int64  `gorm:"column:PK"`
	Code          string `gorm:"column:Code"`
	CodeLowerCase string `gorm:"column:codeLowerCase"`
	ExtensionName string `gorm:"column:p_extensionname"`
}

func (EnumerationEntity) TableName() string {
	return "enumerationvalues"
}

func GetEnumerationByID(id int64) (*EnumerationEntity, error) {
	var enumeration *EnumerationEntity
	DB.Find(&enumeration, "PK = ?", id)

	return enumeration, nil
}
