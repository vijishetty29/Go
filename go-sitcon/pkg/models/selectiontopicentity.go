package models

import (
	"time"
)

type SelectionTopicEntity struct {
	ID             int64           `gorm:"column:PK"`
	Code           string          `gorm:"column:p_code"`
	Status         int64           `gorm:"column:p_status"`
	TopicNr        string          `gorm:"column:p_topicnr"`
	AdvertisingDay string          `gorm:"column:p_advertisingday"`
	SelectionID    int64           `gorm:"column:p_selection"`
	CreatedTS      *time.Time      `json:"createdTS"`
	ModifiedTS     *time.Time      `json:"modifiedTS"`
	Selection      SelectionEntity `gorm:"foreignKey:SelectionID"`
}

func (SelectionTopicEntity) TableName() string {
	return "pimselectiontopic"
}

func GetModifiedSelectionTopics(since time.Time, activeSelectionInts []int64) ([]*SelectionTopicEntity, error) {
	var selectionTopics []*SelectionTopicEntity

	DB.Preload("Selection").Find(&selectionTopics, "ModifiedTS > ? AND p_selection IN ?", since, activeSelectionInts)

	/*
		DB.Table("pimselectiontopic AS topics").
			Select("topics.*, selections.pk as p_selection").
			Joins("JOIN pimselection AS selections ON topics.p_selection = selections.PK").
			Find(&selectionTopics, "topics.ModifiedTS > ? AND topics.p_selection IN ?", since, activeSelectionInts)
	*/

	return selectionTopics, nil
}
