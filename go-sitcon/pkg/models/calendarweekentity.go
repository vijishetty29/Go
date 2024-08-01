package models

type CalendarWeekEntity struct {
	PK             int64  `gorm:"column:PK"`
	ID             int    `gorm:"column:p_id"`
	Week           string `gorm:"column:p_week"`
	Year           string `gorm:"column:p_year"`
	SelectionTopic int64  `gorm:"column:p_selectiontopic"`
	SelectionItem  int64  `gorm:"column:p_selectionitem"`
}

func (CalendarWeekEntity) TableName() string {
	return "pimcalendarweek"
}

func GetCalendarWeeksForSelectionTopic(topicPk int64) []*CalendarWeekEntity {
	var calendarWeeks []*CalendarWeekEntity
	result := DB.Find(&calendarWeeks, "p_selectiontopic = ?", topicPk)

	if result.RowsAffected == 0 {
		return nil
	}

	return calendarWeeks
}

func GetCalendarWeeksForSelectionItem(itemPk int64) []*CalendarWeekEntity {
	var calendarWeeks []*CalendarWeekEntity
	result := DB.Find(&calendarWeeks, "p_selectionitem = ?", itemPk)

	if result.RowsAffected == 0 {
		return nil
	}

	return calendarWeeks
}
