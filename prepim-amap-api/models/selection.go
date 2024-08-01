package models

import (
	"time"
)

type Selection struct {
	PK                   int64      `gorm:"index:column:PK"`
	Code                 string     `gorm:"column:p_code"`
	AdvertisingStartWeek string     `gorm:"column:p_advertisingstartweek"`
	AdvertisingStartYear string     `gorm:"column:p_advertisingstartyear"`
	AdvertisingEndWeek   string     `gorm:"column:p_advertisingendweek"`
	AdvertisingEndYear   string     `gorm:"column:p_advertisingendyear"`
	DeliveryStartWeek    string     `gorm:"column:p_deliverystartweek"`
	DeliveryStartYear    string     `gorm:"column:p_deliverystartyear"`
	DeliveryEndWeek      string     `gorm:"column:p_deliveryendweek"`
	DeliveryEndYear      string     `gorm:"column:p_deliveryendyear"`
	Active               string     `gorm:"column:p_active"`
	CreatedTS            *time.Time `json:"createdTS"`
	ModifiedTS           *time.Time `json:"modifiedTS"`
}

func (Selection) TableName() string {
	return "pimselection"
}

func GetModifiedSelections(since time.Time, activeSelectionInts []int64) ([]*Selection, error) {
	var selections []*Selection
	DB.Find(&selections, "ModifiedTS > ? AND PK IN ?", since, activeSelectionInts)

	return selections, nil
}

func GetActiveSelections() ([]int64, error) {
	var activeSelections []*Selection
	DB.Find(&activeSelections, "p_active = ?", 1)

	activeSelectionInts := []int64{}
	for _, selection := range activeSelections {
		activeSelectionInts = append(activeSelectionInts, selection.PK)
	}
	return activeSelectionInts, nil
}
