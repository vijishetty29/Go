package models

type SelectionTopicLocalizationEntity struct {
	ID          int64  `gorm:"column:ITEMPK"`
	LangPK      int64  `gorm:"column:LANGPK"`
	Description string `gorm:"column:p_description"`
}

func (SelectionTopicLocalizationEntity) TableName() string {
	return "pimselectiontopiclp"
}

func GetDescriptionForSelectionTopicAndLanguage(topicPk int64, langPk string) string {
	var localizedData SelectionTopicLocalizationEntity

	result := DB.First(&localizedData, "ITEMPK = ? AND LANGPK = ?", topicPk, langPk)

	if result.Error == nil {
		return localizedData.Description
	}
	return ""
}
