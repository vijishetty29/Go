package models

import (
	"time"

	"gorm.io/gorm"
)

type SelectionTopic struct {
	ID             int64      `gorm:"column:PK"`
	Code           string     `gorm:"column:p_code"`
	StatusPK       string     `gorm:"column:p_status"`
	Status         Status     `gorm:"foreignKey:StatusPK"`
	TopicNr        string     `gorm:"column:p_topicnr"`
	AdvertisingDay string     `gorm:"column:p_advertisingday"`
	SelectionPK    int64      `gorm:"column:p_selection"`
	Selection      Selection  `gorm:"foreignKey:SelectionPK"`
	CreatedTS      *time.Time `json:"createdTS"`
	ModifiedTS     *time.Time `json:"modifiedTS"`
}

func (SelectionTopic) TableName() string {
	return "pimselectiontopic"
}

func GetModifiedSelectionTopics(since time.Time, db *gorm.DB) ([]*SelectionTopic, error) {
	var selectionTopics []*SelectionTopic
	db.Preload("Selection").Find(&selectionTopics, "ModifiedTS > ?", since)
	return selectionTopics, nil
}
