package models

type ProductLocalizationEntity struct {
	ID               int64  `gorm:"column:ITEMPK"`
	LangPK           int64  `gorm:"column:LANGPK"`
	Name             string `gorm:"column:p_name"`
	Description      string `gorm:"column:p_description"`
	RawMarketingText string `gorm:"column:p_pimrawmarketingtext"`
}

func (ProductLocalizationEntity) TableName() string {
	return "productslp"
}

func GetProductLocalizedEntityForProductAndLanguage(productPk int64, langPk string) *ProductLocalizationEntity {
	var localizedData *ProductLocalizationEntity
	DB.Limit(1).Find(&localizedData, "ITEMPK = ? AND LANGPK = ?", productPk, langPk)
	return localizedData
}
