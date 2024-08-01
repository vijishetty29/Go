package models

import (
	"time"
)

type SelectionEntity struct {
	PK                   int64      `gorm:"primaryKey:index:column:PK"`
	Code                 string     `gorm:"column:p_code"`
	AdvertisingStartWeek int        `gorm:"column:p_advertisingstartweek"`
	AdvertisingStartYear int        `gorm:"column:p_advertisingstartyear"`
	AdvertisingEndWeek   int        `gorm:"column:p_advertisingendweek"`
	AdvertisingEndYear   int        `gorm:"column:p_advertisingendyear"`
	DeliveryStartWeek    int        `gorm:"column:p_deliverystartweek"`
	DeliveryStartYear    int        `gorm:"column:p_deliverystartyear"`
	DeliveryEndWeek      int        `gorm:"column:p_deliveryendweek"`
	DeliveryEndYear      int        `gorm:"column:p_deliveryendyear"`
	Active               int        `gorm:"column:p_active"`
	CreatedTS            *time.Time `json:"createdTS"`
	ModifiedTS           *time.Time `json:"modifiedTS"`
}

func (SelectionEntity) TableName() string {
	return "pimselection"
}

func GetModifiedSelections(since time.Time, activeSelectionInts []int64) ([]*SelectionEntity, error) {
	var selections []*SelectionEntity
	DB.Find(&selections, "ModifiedTS > ? AND PK IN ?", since, activeSelectionInts)

	return selections, nil
}

func GetActiveSelections() ([]int64, error) {
	var activeSelections []*SelectionEntity
	DB.Find(&activeSelections, "p_active = ?", 1)

	activeSelectionInts := []int64{}
	for _, selection := range activeSelections {
		activeSelectionInts = append(activeSelectionInts, selection.PK)
	}
	return activeSelectionInts, nil
}
